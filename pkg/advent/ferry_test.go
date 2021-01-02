package advent

import (
	"strings"
	"testing"
)

func TestNewFerry(t *testing.T) {
	f := NewFerry()

	got := f.ManhattanDistance()
	if got != 0 {
		t.Errorf("ManhattanDistance: got %#v, want %#v", got, 0)
	}
}

func TestSimpleNavigator(t *testing.T) {
	f := NewFerry()

	in := `F10
N3
F7
R90
F11
`

	err := f.Navigate(strings.NewReader(in), SimpleNavigator())
	if err != nil {
		t.Fatal(err)
	}

	got := f.ManhattanDistance()
	if got != 25 {
		t.Errorf("ManhattanDistance: got %#v, want %#v", got, 25)
	}
}

func TestWaypointNavigator(t *testing.T) {
	f := NewFerry()

	in := `F10
N3
F7
R90
F11
`

	err := f.Navigate(strings.NewReader(in), WaypointNavigator(10, 1))
	if err != nil {
		t.Fatal(err)
	}

	got := f.ManhattanDistance()
	if got != 286 {
		t.Errorf("ManhattanDistance: got %#v, want %#v", got, 286)
	}
}
