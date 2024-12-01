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
		inputPath string
		want      int
	}{
		{inputPath: "testdata/example1.txt", want: 11},
		{inputPath: "testdata/input.txt", want: 1651298},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Part 1: %s", test.inputPath), func(t *testing.T) {
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

			totalDistance := 0

			for i, id := range lists[0] {
				distance := id - lists[1][i]
				if distance < 0 {
					totalDistance -= distance
				} else {
					totalDistance += distance
				}
			}

			if totalDistance != test.want {
				t.Errorf("totalDistance: got %d, want %d", totalDistance, test.want)
			}
		})
	}
}
