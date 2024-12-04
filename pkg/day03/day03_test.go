package day03

import (
	"fmt"
	"testing"
)

func TestDay03(t *testing.T) {
	tests := []struct {
		inputPath            string
		wantResult           int
		wantWithConditionals int
	}{
		{inputPath: "testdata/example1.txt", wantResult: 161, wantWithConditionals: 161},
		{inputPath: "testdata/example2.txt", wantResult: 161, wantWithConditionals: 48},
		{inputPath: "testdata/input.txt", wantResult: 189600467, wantWithConditionals: 107069718},
	}

	for _, test := range tests {
		input, err := load(test.inputPath)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Part 1: %s", test.inputPath), func(t *testing.T) {
			result := input.process()

			if result != test.wantResult {
				t.Errorf("process: got %d, want %d", result, test.wantResult)
			}
		})

		t.Run(fmt.Sprintf("Part 2: %s", test.inputPath), func(t *testing.T) {
			result := input.process(DO, DONT)

			if result != test.wantWithConditionals {
				t.Errorf("process (with conditionals): got %d, want %d", result, test.wantWithConditionals)
			}
		})
	}
}
