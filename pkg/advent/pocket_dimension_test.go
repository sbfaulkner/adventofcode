package advent

import (
	"strings"
	"testing"
)

func TestPocketDimension3(t *testing.T) {
	in := `.#.
..#
###
`

	d, err := ReadPocketDimension(strings.NewReader(in))
	if err != nil {
		t.Error(err)
	}

	t.Run("Cycle 1", func(t *testing.T) {
		got := d.Cycle(3)
		if got != 11 {
			t.Errorf("got %d, want %d", got, 11)
		}
	})

	t.Run("Cycle 2", func(t *testing.T) {
		got := d.Cycle(3)
		if got != 21 {
			t.Errorf("got %d, want %d", got, 21)
		}
	})

	t.Run("Cycle 3", func(t *testing.T) {
		got := d.Cycle(3)
		if got != 38 {
			t.Errorf("got %d, want %d", got, 38)
		}
	})

	// Cycle 4
	d.Cycle(3)

	// Cycle 5
	d.Cycle(3)

	t.Run("Cycle 6", func(t *testing.T) {
		got := d.Cycle(3)
		if got != 112 {
			t.Errorf("got %d, want %d", got, 112)
		}
	})
}

func TestPocketDimension4(t *testing.T) {
	in := `.#.
..#
###
`

	d, err := ReadPocketDimension(strings.NewReader(in))
	if err != nil {
		t.Error(err)
	}

	t.Run("Cycle 1", func(t *testing.T) {
		got := d.Cycle(4)
		if got != 29 {
			t.Errorf("got %d, want %d", got, 29)
		}
	})

	t.Run("Cycle 2", func(t *testing.T) {
		got := d.Cycle(4)
		if got != 60 {
			t.Errorf("got %d, want %d", got, 60)
		}
	})

	// Cycle 3
	d.Cycle(4)

	// Cycle 4
	d.Cycle(4)

	// Cycle 5
	d.Cycle(4)

	t.Run("Cycle 6", func(t *testing.T) {
		got := d.Cycle(4)
		if got != 848 {
			t.Errorf("got %d, want %d", got, 848)
		}
	})
}
