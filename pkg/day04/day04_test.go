package day04

import (
	"testing"
)

func TestDay04Part1Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 18
	got := input.part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDay04Part1(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 2646
	got := input.part1()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDay04Part2Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 9
	got := input.part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDay04Part2(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 2000
	got := input.part2()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
