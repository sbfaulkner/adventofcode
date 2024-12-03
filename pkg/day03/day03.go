package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type input struct {
	args [][]int
}

var reMUL = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func scanMatches(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0

	loc := reMUL.FindIndex(data[start:])
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
		args: make([][]int, 0),
	}

	for scanner.Scan() {
		text := scanner.Text()

		matches := reMUL.FindStringSubmatch(text)

		s := matches[1:]
		args := make([]int, len(s))

		for i, v := range s {
			args[i], err = strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("failed to convert operand %s to int: %w", v, err)
			}
		}

		input.args = append(input.args, args)
	}

	return &input, nil
}

func (i *input) sumOfProducts() int {
	sum := 0

	for _, args := range i.args {
		sum += args[0] * args[1]
	}

	return sum
}
