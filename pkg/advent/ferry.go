package advent

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

// Ferry position and orientation
type Ferry struct {
	dir  int
	x, y int
}

const (
	moveNorth   = 'N'
	moveSouth   = 'S'
	moveEast    = 'E'
	moveWest    = 'W'
	turnLeft    = 'L'
	turnRight   = 'R'
	moveForward = 'F'
)

const (
	east  = 0
	north = 90
	west  = 180
	south = 270
)

// NewFerry creates a new ferry pointing in the right direction with origin (0, 0)
func NewFerry() *Ferry {
	return &Ferry{}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

// ManhattanDistance returns Manhattan distance from its starting position
func (f Ferry) ManhattanDistance() int {
	return abs(f.x) + abs(f.y)
}

// Navigator function that interprets an action-value pair to move a Ferry
type Navigator func(f *Ferry, action rune, value int)

// Navigate moves the ferry using commands from the reader
func (f *Ferry) Navigate(rd io.Reader, navigate Navigator) error {
	s := bufio.NewScanner(rd)

	var action rune
	var value int

	for s.Scan() {
		fmt.Sscanf(s.Text(), "%c%d", &action, &value)
		navigate(f, action, value)
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}

// SimpleNavigator implements basic ferry-relative navigation
func SimpleNavigator() Navigator {
	return func(f *Ferry, action rune, value int) {
		switch action {
		case moveNorth:
			f.x, f.y = move(f.x, f.y, north, value)
		case moveSouth:
			f.x, f.y = move(f.x, f.y, south, value)
		case moveEast:
			f.x, f.y = move(f.x, f.y, east, value)
		case moveWest:
			f.x, f.y = move(f.x, f.y, west, value)
		case turnLeft:
			f.dir += value
		case turnRight:
			f.dir -= value
		case moveForward:
			f.x, f.y = move(f.x, f.y, f.dir, value)
		}
	}
}

// WaypointNavigator implements waypoint-relativer navigation
func WaypointNavigator(x, y int) Navigator {
	return func(f *Ferry, action rune, value int) {
		switch action {
		case moveNorth:
			x, y = move(x, y, north, value)
		case moveSouth:
			x, y = move(x, y, south, value)
		case moveEast:
			x, y = move(x, y, east, value)
		case moveWest:
			x, y = move(x, y, west, value)
		case turnLeft:
			x, y = rotate(x, y, value)
		case turnRight:
			x, y = rotate(x, y, -value)
		case moveForward:
			f.x += x * value
			f.y += y * value
		}
	}
}

func move(x, y int, angle int, units int) (int, int) {
	sin, cos := math.Sincos(float64(angle) / 180 * math.Pi)

	x += int(cos * float64(units))
	y += int(sin * float64(units))

	return x, y
}

func rotate(x, y int, angle int) (int, int) {
	sin, cos := math.Sincos(float64(angle) / 180 * math.Pi)

	px := float64(x)
	py := float64(y)

	fx := math.Round(px*cos - py*sin)
	fy := math.Round(px*sin + py*cos)

	return int(fx), int(fy)
}
