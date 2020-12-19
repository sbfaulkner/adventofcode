package tree

import (
	"strings"
	"testing"
)

func readMap(t *testing.T) *Map {
	in := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

	m, err := ReadMap(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	return m
}

func TestCount(t *testing.T) {
	m := readMap(t)

	testCases := []struct {
		s    Slope
		want int
	}{
		{Slope{3, 1}, 7},
	}

	for _, tc := range testCases {
		got := m.Count(tc.s)
		if got != tc.want {
			t.Errorf("m.Count(%#v): got %#v, want %#v", tc.s, got, tc.want)
		}
	}
}

func TestProductOfCounts(t *testing.T) {
	m := readMap(t)

	testCases := []struct {
		s    []Slope
		want int
	}{
		{
			[]Slope{
				{1, 1},
				{3, 1},
				{5, 1},
				{7, 1},
				{1, 2},
			},
			336,
		},
	}

	for _, tc := range testCases {
		got := m.ProductOfCounts(tc.s)

		if got != tc.want {
			t.Errorf("m.ProductOfCounts(): got %v, want %v", got, tc.want)
		}
	}
}
