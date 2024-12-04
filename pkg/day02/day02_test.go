package day02

import (
	"testing"
)

func TestDay02Part1Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 2
	got := input.safe()

	if got != want {
		t.Errorf("safe: got %d, want %d", got, want)
	}
}

func TestDay02Part1(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 534
	got := input.safe()

	if got != want {
		t.Errorf("safe: got %d, want %d", got, want)
	}
}

func TestDay02Part2Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 4
	got := input.safeWithDampener()

	if got != want {
		t.Errorf("safeWithDampener: got %d, want %d", got, want)
	}
}

func TestDay02Part2(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 577
	got := input.safeWithDampener()

	if got != want {
		t.Errorf("safeWithDampener: got %d, want %d", got, want)
	}
}
