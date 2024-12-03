package day03

import (
	"fmt"
	"testing"
)

func TestDay03(t *testing.T) {
	tests := []struct {
		inputPath         string
		wantSumOfProducts int
	}{
		{inputPath: "testdata/example1.txt", wantSumOfProducts: 161},
		{inputPath: "testdata/input.txt", wantSumOfProducts: 189600467},
	}

	for _, test := range tests {
		input, err := load(test.inputPath)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Part 1: %s", test.inputPath), func(t *testing.T) {
			sumOfProducts := input.sumOfProducts()

			if sumOfProducts != test.wantSumOfProducts {
				t.Errorf("sumOfProducts: got %d, want %d", sumOfProducts, test.wantSumOfProducts)
			}
		})
	}
}
