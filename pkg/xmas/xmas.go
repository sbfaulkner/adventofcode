package xmas

import (
	"bufio"
	"io"
	"strconv"
)

// XMAS represents a stream of data encoded using eXchange-Masking Addition System (XMAS)
type XMAS struct {
	data []int
	off  int
	size int
}

// LoadXMAS reads an eXchange-Masking Addition System (XMAS) encoded data stream
func LoadXMAS(size int, rd io.Reader) (*XMAS, error) {
	xmas := XMAS{
		data: make([]int, 0, size),
		off:  size,
		size: size,
	}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}

		xmas.data = append(xmas.data, n)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return &xmas, nil
}

// FindInvalid finds the first value in the data stream that is not valid
func (x *XMAS) FindInvalid() int {
	for x.off < len(x.data) {
		if !x.isValid(x.data[x.off]) {
			return x.data[x.off]
		}

		x.off++
	}

	return 0
}

func (x XMAS) isValid(number int) bool {
	for i, m := range x.data[x.off-x.size : x.off-1] {
		for _, n := range x.data[x.off-x.size+i+1 : x.off] {
			if number == m+n {
				return true
			}
		}
	}

	return false
}

// FindRangeMinMax finds the min and max of the contiguous set of numbers that add up to a given value
func (x XMAS) FindRangeMinMax(sum int) (int, int) {
	for i := range x.data[:len(x.data)-1] {
		for j := range x.data[i+1:] {
			if sumEqual(sum, x.data[i:i+1+j+1]) {
				return minMax(x.data[i : i+1+j+1])
			}
		}
	}

	return 0, 0
}

func sumEqual(sum int, series []int) bool {
	s := 0

	for _, n := range series {
		s += n

		if s > sum {
			return false
		}
	}

	return s == sum
}

func minMax(series []int) (int, int) {
	min := series[0]
	max := series[0]

	for _, n := range series[1:] {
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}

	return min, max
}
