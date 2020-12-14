package main

import (
	"strings"
	"testing"
)

func TestTreeCount(t *testing.T) {
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
	slopes := []slope{
		{3, 1},
	}

	want := 7

	got := TreeCount(strings.NewReader(in), slopes)

	if got[0] != want {
		t.Errorf("TreeCount: got %v, want %v", got, want)
	}
}

func TestTreeCountProduct(t *testing.T) {
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
	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	want := 336

	got := TreeCountProduct(strings.NewReader(in), slopes)

	if got != want {
		t.Errorf("TreeCountProduct: got %v, want %v", got, want)
	}
}
