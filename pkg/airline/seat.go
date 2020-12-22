package airline

import (
	"bufio"
	"io"
)

// Seat for airline
type Seat struct {
	Row    int
	Column int
	ID     int
}

// NewSeat creats a new Seat
func NewSeat(input string) *Seat {
	r := 0
	rows := 128
	c := 0
	columns := 8

	for i, ch := range []rune(input) {
		if i < 7 {
			rows /= 2
			if ch == 'B' {
				r += rows
			}
		} else {
			columns /= 2
			if ch == 'R' {
				c += columns
			}
		}
	}

	return &Seat{
		Row:    r,
		Column: c,
		ID:     r*8 + c,
	}
}

// ReadSeats reads all seat codes from the provided reader
func ReadSeats(rd io.Reader) ([]*Seat, error) {
	seats := []*Seat{}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		seats = append(seats, NewSeat(s.Text()))
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}
