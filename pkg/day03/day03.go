package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	MUL = iota
	DO
	DONT
)

var INSTRUCTIONS = map[string]int{
	"do":    DO,
	"don't": DONT,
	"mul":   MUL,
}

type input struct {
	instructions []int
}

var reINSTRUCTION = regexp.MustCompile(`(do|don't|mul)\((?:(\d+),(\d+))?\)`)

func scanMatches(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0

	loc := reINSTRUCTION.FindIndex(data[start:])
	if loc == nil {
		return start, nil, nil
	}

	return start + loc[1], data[start+loc[0] : start+loc[1]], nil
}

func load(inputPath string) (*input, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(scanMatches)

	input := input{
		instructions: make([]int, 0),
	}

	for scanner.Scan() {
		text := scanner.Text()

		matches := reINSTRUCTION.FindStringSubmatch(text)

		instructions := make([]int, 1, len(matches)-1)
		instructions[0] = INSTRUCTIONS[matches[1]]

		for _, m := range matches[2:] {
			if m == "" {
				break
			}

			a, err := strconv.Atoi(m)
			if err != nil {
				return nil, fmt.Errorf("failed to convert operand %s to int: %w", m, err)
			}

			instructions = append(instructions, a)
		}

		input.instructions = append(input.instructions, instructions...)
	}

	return &input, nil
}

func (i *input) process(extensions ...int) int {
	enabled := true
	sum := 0

	extensionMap := make(map[int]bool, len(extensions)+1)
	extensionMap[MUL] = true

	for _, e := range extensions {
		extensionMap[e] = true
	}

	for pc := 0; pc < len(i.instructions); pc++ {
		if !extensionMap[i.instructions[pc]] {
			continue
		}

		switch i.instructions[pc] {
		case DO:
			enabled = true
		case DONT:
			enabled = false
		case MUL:
			if enabled {
				sum += i.instructions[pc+1] * i.instructions[pc+2]
			}

			pc += 2
		}
	}

	return sum
}
