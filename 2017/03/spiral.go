package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(value int) int {
	if value < 0 {
		value = -value
	}

	return value
}

type spiral struct {
	last int
	size int
	grid [][]int
}

func newSpiral(last int) spiral {
	size := int(math.Ceil(math.Sqrt(float64(last))))

	s := spiral{last: last, size: size}

	s.grid = make([][]int, size)

	for i := range s.grid {
		s.grid[i] = make([]int, size)
	}

	return s
}

func (s spiral) fill() {
	for i := 1; i <= s.last; i++ {
		x, y := s.position(i)
		s.grid[y][x] = i
	}
}

func (s spiral) middle() (x, y int) {
	x = (s.size - 1) / 2
	y = s.size / 2

	return
}

func (s spiral) position(value int) (x, y int) {
	square := s.size * s.size
	diff := square - value

	if s.size%2 == 1 {
		x = s.size - 1
		y = s.size - 1

		if diff >= s.size {
			x = x - (s.size - 1)
		}

		y = y - diff%s.size
	} else {
		x = 0
		y = 0

		if diff >= s.size {
			x = x + (s.size - 1)
		}

		y = y + diff%s.size
	}

	return
}

func (s spiral) distance() int {
	x1, y1 := s.middle()
	x2, y2 := s.position(s.last)

	return abs(x2-x1) + abs(y2-y1)
}

func main() {
	input := 289326

	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		check(err)
		input = i
	}

	s := newSpiral(input)

	fmt.Println(s.distance())
}
