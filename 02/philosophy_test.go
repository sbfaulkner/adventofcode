package main

import (
	"strings"
	"testing"
)

func TestCountValidSledPolicy(t *testing.T) {
	in := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

	want := 2

	got := CountValid(strings.NewReader(in), sledPolicy)

	if got != want {
		t.Errorf("CountValid: got %v, want %v", got, want)
	}
}

func TestCountValidTobogganPolicy(t *testing.T) {
	in := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

	want := 1

	got := CountValid(strings.NewReader(in), tobogganPolicy)

	if got != want {
		t.Errorf("CountValid: got %v, want %v", got, want)
	}
}
