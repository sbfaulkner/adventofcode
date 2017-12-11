package main

import (
	"fmt"
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

type direction int

const (
	east direction = iota
	north
	west
	south
)

type spiral struct {
	last int
	x, y int
	d    direction
	grid [][]int
}

func newSpiral(last int) spiral {
	s := spiral{d: east, last: last, grid: [][]int{[]int{1}}}

	s.fill()

	return s
}

func (s spiral) value() int {
	return s.grid[s.y][s.x]
}

func (s *spiral) fill() {
	for s.value() < s.last {
		s.setNext()
	}
}

func (s *spiral) addColumn() {
	for i := range s.grid {
		s.grid[i] = append(s.grid[i], 0)
	}
}

func (s *spiral) insertColumn() {
	for i := range s.grid {
		s.grid[i] = append([]int{0}, s.grid[i]...)
	}
}

func (s *spiral) addRow() {
	s.grid = append(s.grid, make([]int, len(s.grid[0])))
}

func (s *spiral) insertRow() {
	s.grid = append([][]int{make([]int, len(s.grid[0]))}, s.grid...)
}

func (s *spiral) setNext() {
	v := s.value()

	switch s.d {
	case east:
		s.x++
		if s.x == len(s.grid[s.y]) {
			s.addColumn()
			s.d = north
		}
	case north:
		if s.y > 0 {
			s.y--
		} else {
			s.insertRow()
			s.d = west
		}
	case west:
		if s.x > 0 {
			s.x--
		} else {
			s.insertColumn()
			s.d = south
		}
	case south:
		s.y++
		if s.y == len(s.grid) {
			s.addRow()
			s.d = east
		}
	}

	v++

	s.grid[s.y][s.x] = v
}

func (s spiral) middle() (x, y int) {
	x = (len(s.grid[0]) - 1) / 2
	y = len(s.grid) / 2

	return
}

func (s spiral) distance() int {
	x, y := s.middle()

	return abs(s.x-x) + abs(s.y-y)
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
