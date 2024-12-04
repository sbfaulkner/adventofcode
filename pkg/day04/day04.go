package day04

import (
	"fmt"
	"os"
)

type input struct{}

func load(inputPath string) (*input, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	// scanner := bufio.NewScanner(file)

	input := input{}

	return &input, nil
}

func (i *input) part1() int {
	return 0
}

func (i *input) part2() int {
	return 0
}
