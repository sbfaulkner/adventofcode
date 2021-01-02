package advent

import (
	"bufio"
	"io"
)

// Map of trees on a hill
type Map [][]rune

// Slope of path down hill
type Slope struct {
	DX, DY int
}

// ReadMap creates a new Map loading it from a Reader
func ReadMap(rd io.Reader) (*Map, error) {
	m := Map{}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		m = append(m, []rune(s.Text()))
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return &m, nil
}

// Count returns the number of trees on a given Slope
func (m Map) Count(s Slope) int {
	count := 0

	var x, y int

	for y < len(m) {
		if m[y][x] == '#' {
			count++
		}

		y += s.DY
		x += s.DX
		x %= len(m[0])
	}

	return count
}

// ProductOfCounts returns the product of the tree counts for the trajectories
func (m Map) ProductOfCounts(slopes []Slope) int {
	p := 1

	for _, s := range slopes {
		p *= m.Count(s)
	}

	return p
}
