package main

import (
	"strings"
	"testing"
)

func TestFix(t *testing.T) {
	in := `1721
979
366
299
675
1456
`

	want := 514579

	got := Fix(strings.NewReader(in))

	if got != want {
		t.Errorf("Fix: got %v, want %v", got, want)
	}
}

func TestFix3(t *testing.T) {
	in := `1721
979
366
299
675
1456
`

	want := 241861950

	got := Fix3(strings.NewReader(in))

	if got != want {
		t.Errorf("Fix3: got %v, want %v", got, want)
	}
}
