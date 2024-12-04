package day03

import (
	"testing"
)

func TestDay03Part1Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 161
	got := input.process()

	if got != want {
		t.Errorf("process: got %d, want %d", got, want)
	}
}

func TestDay03Part1Example2(t *testing.T) {
	input, err := load("testdata/example2.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 161
	got := input.process()

	if got != want {
		t.Errorf("process: got %d, want %d", got, want)
	}
}

func TestDay03Part1(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 189600467
	got := input.process()

	if got != want {
		t.Errorf("process: got %d, want %d", got, want)
	}
}

func TestDay03Part2Example1(t *testing.T) {
	input, err := load("testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 161
	got := input.process(DO, DONT)

	if got != want {
		t.Errorf("process (with conditionals): got %d, want %d", got, want)
	}
}

func TestDay03Part2Example2(t *testing.T) {
	input, err := load("testdata/example2.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 48
	got := input.process(DO, DONT)

	if got != want {
		t.Errorf("process (with conditionals): got %d, want %d", got, want)
	}
}

func TestDay03Part2(t *testing.T) {
	input, err := load("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 107069718
	got := input.process(DO, DONT)

	if got != want {
		t.Errorf("process (with conditionals): got %d, want %d", got, want)
	}
}
