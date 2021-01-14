package advent

import (
	"bufio"
	"io"
)

// PocketDimension containing Conway cubes
type PocketDimension struct {
	space [][][][]rune
}

const (
	activeCube   = '#'
	inactiveCube = '.'
)

// ReadPocketDimension reads an initial slice of cubes from a reader
func ReadPocketDimension(rd io.Reader) (*PocketDimension, error) {
	s := bufio.NewScanner(rd)

	d := PocketDimension{}

	d.space = make([][][][]rune, 1)
	d.space[0] = make([][][]rune, 1)

	for s.Scan() {
		d.space[0][0] = append(d.space[0][0], []rune(s.Text()))
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return &d, nil
}

func minmax(i, imin, imax int) (int, int) {
	return min(i, imin), max(i, imax)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// Cycle updates the state of the dimension
func (d *PocketDimension) Cycle(dimensions int) int {
	var xmin, ymin, zmin, wmin, xmax, ymax, zmax, wmax int

	count := 0

	winc := 0
	if dimensions > 3 {
		winc = 1
	}

	wsize := len(d.space) + 2*winc
	zsize := len(d.space[0]) + 2
	ysize := len(d.space[0][0]) + 2
	xsize := len(d.space[0][0][0]) + 2

	wmin = wsize - 1
	zmin = zsize - 1
	ymin = ysize - 1
	xmin = xsize - 1

	cycle := make([][][][]rune, wsize)

	for w := 0; w < wsize; w++ {
		cycle[w] = make([][][]rune, zsize)
		for z := 0; z < zsize; z++ {
			cycle[w][z] = make([][]rune, ysize)
			for y := 0; y < ysize; y++ {
				cycle[w][z][y] = make([]rune, xsize)
				for x := 0; x < xsize; x++ {
					a := d.activeNeighbours(x-1, y-1, z-1, w-winc)
					if d.isActive(x-1, y-1, z-1, w-winc) && a == 2 || a == 3 {
						cycle[w][z][y][x] = activeCube

						count++

						wmin, wmax = minmax(w, wmin, wmax)
						zmin, zmax = minmax(z, zmin, zmax)
						ymin, ymax = minmax(y, ymin, ymax)
						xmin, xmax = minmax(x, xmin, xmax)
					} else {
						cycle[w][z][y][x] = inactiveCube
					}
				}
			}
		}
	}

	wsize = wmax - wmin + 1
	zsize = zmax - zmin + 1
	ysize = ymax - ymin + 1
	xsize = xmax - xmin + 1

	d.space = make([][][][]rune, wsize)
	for w := 0; w < wsize; w++ {
		d.space[w] = make([][][]rune, zsize)
		for z := 0; z < zsize; z++ {
			d.space[w][z] = make([][]rune, ysize)
			for y := 0; y < ysize; y++ {
				d.space[w][z][y] = cycle[w+wmin][z+zmin][y+ymin][xmin : xmax+1]
			}
		}
	}

	return count
}

func (d PocketDimension) activeNeighbours(x, y, z, w int) int {
	count := 0

	for dw := -1; dw <= 1; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dw == 0 && dz == 0 && dy == 0 && dx == 0 {
						continue
					}
					if d.isActive(x+dx, y+dy, z+dz, w+dw) {
						count++
					}
				}
			}
		}
	}

	return count
}

func (d PocketDimension) isActive(x, y, z, w int) bool {
	if w < 0 || w >= len(d.space) {
		return false
	}

	if z < 0 || z >= len(d.space[w]) {
		return false
	}

	if y < 0 || y >= len(d.space[w][z]) {
		return false
	}

	if x < 0 || x >= len(d.space[w][z][y]) {
		return false
	}

	return d.space[w][z][y][x] == activeCube
}
