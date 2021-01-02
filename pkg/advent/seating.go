package advent

import (
	"bufio"
	"io"
)

const (
	floor        = '.'
	emptySeat    = 'L'
	occupiedSeat = '#'
)

// Seating layout of seats (and floor)
type Seating [][]rune

// ReadSeating reads a ferry floorplan from the provided reader
func ReadSeating(rd io.Reader) (*Seating, error) {
	seating := Seating{}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		seating = append(seating, []rune(s.Text()))
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return &seating, nil
}

// Evolve will update the ferry seating
func (s Seating) Evolve(distance int, tolerance int) (*Seating, int, bool) {
	e := make(Seating, 0, len(s))

	occupied := 0
	evolved := false

	for y, row := range s {
		er := make([]rune, 0, len(row))

		for x, pos := range row {
			switch pos {
			case emptySeat:
				if s.countVisibileFrom(x, y, distance) == 0 {
					er = append(er, occupiedSeat)
					evolved = true
					occupied++
				} else {
					er = append(er, emptySeat)
				}
			case occupiedSeat:
				if s.countVisibileFrom(x, y, distance) >= tolerance {
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

func (s Seating) countVisibileFrom(x, y int, distance int) int {
	n := 0

	for _, v := range vectors {
		for d := 1; d <= distance; d++ {
			py := y + d*v.dy
			if py < 0 || py >= len(s) {
				break
			}

			px := x + d*v.dx
			if px < 0 || px >= len(s[py]) {
				break
			}

			p := s[py][px]
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
