package day05

import (
	"testing"
)

func TestDay05Part1Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 143
	got := input.part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDay05Part1(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 6498
	got := input.part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDay05Part2Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 123
	got := input.part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDay05Part2(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 5017
	got := input.part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
