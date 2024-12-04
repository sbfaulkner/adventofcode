package day04

import (
	"testing"
)

func TestDay04Part1Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 161
	got := input.part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// func TestDay04Part1(t *testing.T) {
// 	input, err := load("testdata/input.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	want := 189600467
// 	got := input.part1()

// 	if got != want {
// 		t.Errorf("got %d, want %d", got, want)
// 	}
// }

// func TestDay04Part2Example2(t *testing.T) {
// 	input, err := load("testdata/example2.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	want := 48
// 	got := input.part2()

// 	if got != want {
// 		t.Errorf("got %d, want %d", got, want)
// 	}
// }

// func TestDay04Part2(t *testing.T) {
// 	input, err := load("testdata/input.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	want := 107069718
// 	got := input.part2()

// 	if got != want {
// 		t.Errorf("got %d, want %d", got, want)
// 	}
// }
