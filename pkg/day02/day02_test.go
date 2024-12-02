package day02

import (
	"fmt"
	"testing"
)

func TestDay02(t *testing.T) {
	tests := []struct {
		inputPath string
		wantSafe  int
	}{
		{inputPath: "testdata/example1.txt", wantSafe: 2},
		{inputPath: "testdata/input.txt", wantSafe: 534},
	}

	for _, test := range tests {
		input, err := load(test.inputPath)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Part 1: %s", test.inputPath), func(t *testing.T) {
			safe := input.safe()

			if safe != test.wantSafe {
				t.Errorf("safe: got %d, want %d", safe, test.wantSafe)
			}
		})
	}
}
