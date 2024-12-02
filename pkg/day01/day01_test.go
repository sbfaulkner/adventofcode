package day01

import (
	"fmt"
	"testing"
)

func TestDay01(t *testing.T) {
	tests := []struct {
		inputPath      string
		wantDistance   int
		wantSimilarity int
	}{
		{inputPath: "testdata/example1.txt", wantDistance: 11, wantSimilarity: 31},
		{inputPath: "testdata/input.txt", wantDistance: 1651298, wantSimilarity: 21306195},
	}

	for _, test := range tests {
		input, err := load(test.inputPath)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Part 1: %s", test.inputPath), func(t *testing.T) {
			distance := input.distance()

			if distance != test.wantDistance {
				t.Errorf("distance: got %d, want %d", distance, test.wantDistance)
			}
		})

		t.Run(fmt.Sprintf("Part 2: %s", test.inputPath), func(t *testing.T) {
			similarity := input.similarity()

			if similarity != test.wantSimilarity {
				t.Errorf("similarity: got %d, want %d", similarity, test.wantSimilarity)
			}
		})
	}
}
