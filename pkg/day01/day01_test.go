package day01

import (
	"testing"
)

func TestDay01Part1Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 11
	got := input.distance()

	if got != want {
		t.Errorf("distance: got %d, want %d", got, want)
	}
}

func TestDay01Part1(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 1651298
	got := input.distance()

	if got != want {
		t.Errorf("distance: got %d, want %d", got, want)
	}
}

func TestDay01Part2Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 31
	got := input.similarity()

	if got != want {
		t.Errorf("similarity: got %d, want %d", got, want)
	}
}

func TestDay01Part2(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 21306195
	got := input.similarity()

	if got != want {
		t.Errorf("similarity: got %d, want %d", got, want)
	}
}
