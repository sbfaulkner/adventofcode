package day01

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type input struct {
	lists [][]int
}

func load(inputPath string) (*input, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	input := input{
		lists: make([][]int, 2),
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		for f, field := range fields {
			id, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("failed to parse field value %s: %w", field, err)
			}

			i, _ := slices.BinarySearch(input.lists[f], id)
			input.lists[f] = slices.Insert(input.lists[f], i, id)
		}
	}

	return &input, nil
}

func (i input) distance() int {
	distance := 0

	for idx, id := range i.lists[0] {
		d := id - i.lists[1][idx]

		if d < 0 {
			distance -= d
		} else {
			distance += d
		}
	}

	return distance
}

func (i input) similarity() int {
	counts := make(map[int]int)

	for _, id := range i.lists[1] {
		counts[id]++
	}

	score := 0

	for _, id := range i.lists[0] {
		score += id * counts[id]
	}

	return score
}
