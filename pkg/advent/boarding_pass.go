package advent

import (
	"bufio"
	"io"
)

// BoardingPass for airline
type BoardingPass struct {
	Row    int
	Column int
	ID     int
}

// NewBoardingPass creats a new BoardingPass
func NewBoardingPass(input string) *BoardingPass {
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

	return &BoardingPass{
		Row:    r,
		Column: c,
		ID:     r*8 + c,
	}
}

// ReadBoardingPasses reads all seat codes from the provided reader
func ReadBoardingPasses(rd io.Reader) ([]*BoardingPass, error) {
	bps := []*BoardingPass{}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		bps = append(bps, NewBoardingPass(s.Text()))
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return bps, nil
}
