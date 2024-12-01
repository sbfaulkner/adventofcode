package day01

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestDay01(t *testing.T) {
	tests := []struct {
		inputPath           string
		wantDistance        int
		wantSimilarityScore int
	}{
		{inputPath: "testdata/example1.txt", wantDistance: 11, wantSimilarityScore: 31},
		{inputPath: "testdata/input.txt", wantDistance: 1651298, wantSimilarityScore: 21306195},
	}

	for _, test := range tests {
		input, err := os.ReadFile(test.inputPath)
		if err != nil {
			t.Fatal(err)
		}

		lists := make([][]int, 2)

		scanner := bufio.NewScanner(bytes.NewReader(input))
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line)

			for f, field := range fields {
				id, err := strconv.Atoi(field)
				if err != nil {
					t.Fatal(err)
				}

				i, _ := slices.BinarySearch(lists[f], id)
				lists[f] = slices.Insert(lists[f], i, id)
			}
		}

		t.Run(fmt.Sprintf("Part 1: %s", test.inputPath), func(t *testing.T) {
			totalDistance := 0

			for i, id := range lists[0] {
				distance := id - lists[1][i]
				if distance < 0 {
					totalDistance -= distance
				} else {
					totalDistance += distance
				}
			}

			if totalDistance != test.wantDistance {
				t.Errorf("totalDistance: got %d, want %d", totalDistance, test.wantDistance)
			}
		})

		t.Run(fmt.Sprintf("Part 2: %s", test.inputPath), func(t *testing.T) {
			locationCounts := make(map[int]int)

			for _, id := range lists[1] {
				locationCounts[id]++
			}

			similarityScore := 0

			for _, id := range lists[0] {
				similarityScore += id * locationCounts[id]
			}

			if similarityScore != test.wantSimilarityScore {
				t.Errorf("similarityScore: got %d, want %d", similarityScore, test.wantSimilarityScore)
			}
		})
	}
}
