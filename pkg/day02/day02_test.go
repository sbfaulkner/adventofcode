package day02

import (
	"fmt"
	"testing"
)

func TestDay02(t *testing.T) {
	tests := []struct {
		inputPath            string
		wantSafe             int
		wantSafeWithDampener int
	}{
		{inputPath: "testdata/example1.txt", wantSafe: 2, wantSafeWithDampener: 4},
		{inputPath: "testdata/input.txt", wantSafe: 534, wantSafeWithDampener: 577},
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

		t.Run(fmt.Sprintf("Part 2: %s", test.inputPath), func(t *testing.T) {
			safe := input.safeWithDampener()

			if safe != test.wantSafeWithDampener {
				t.Errorf("safeWithDampener: got %d, want %d", safe, test.wantSafeWithDampener)
			}
		})
	}
}
