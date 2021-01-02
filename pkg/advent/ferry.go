package advent

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	floor        = '.'
	emptySeat    = 'L'
	occupiedSeat = '#'
)

// Ferry layout of seats (and floor)
type Ferry [][]rune

// ReadFerry reads a ferry floorplan from the provided reader
func ReadFerry(rd io.Reader) (*Ferry, error) {
	ferry := Ferry{}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		ferry = append(ferry, []rune(s.Text()))
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return &ferry, nil
}

func (f Ferry) String() string {
	s := make([]string, 0, len(f))

	for _, r := range f {
		s = append(s, fmt.Sprintf("%q", string(r)))
	}

	return strings.Join(s, ",")
}

// Evolve will update the ferry seating
func (f Ferry) Evolve(distance int, tolerance int) (*Ferry, int, bool) {
	e := make(Ferry, 0, len(f))

	occupied := 0
	evolved := false

	for y, row := range f {
		er := make([]rune, 0, len(row))

		for x, pos := range row {
			switch pos {
			case emptySeat:
				if f.countVisibileFrom(x, y, distance) == 0 {
					er = append(er, occupiedSeat)
					evolved = true
					occupied++
				} else {
					er = append(er, emptySeat)
				}
			case occupiedSeat:
				if f.countVisibileFrom(x, y, distance) >= tolerance {
					er = append(er, emptySeat)
					evolved = true
				} else {
					occupied++
					er = append(er, occupiedSeat)
				}
			default:
				er = append(er, pos)
			}
		}

		e = append(e, er)
	}

	return &e, occupied, evolved
}

var vectors = []struct{ dx, dy int }{
	{dx: -1, dy: -1},
	{dx: 0, dy: -1},
	{dx: +1, dy: -1},
	{dx: -1, dy: 0},
	{dx: +1, dy: 0},
	{dx: -1, dy: +1},
	{dx: 0, dy: +1},
	{dx: +1, dy: +1},
}

func (f Ferry) countVisibileFrom(x, y int, distance int) int {
	n := 0

	for _, v := range vectors {
		for d := 1; d <= distance; d++ {
			py := y + d*v.dy
			if py < 0 || py >= len(f) {
				break
			}

			px := x + d*v.dx
			if px < 0 || px >= len(f[py]) {
				break
			}

			p := f[py][px]
			if p == emptySeat {
				break
			}
			if p == occupiedSeat {
				n++
				break
			}
		}
	}

	return n
}

func (f Ferry) isOccupiedAt(x, y int) bool {
	if y < 0 || y >= len(f) {
		return false
	}

	if x < 0 || x >= len(f[y]) {
		return false
	}

	return f[y][x] == occupiedSeat
}
