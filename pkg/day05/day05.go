package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type set map[int]bool

type input struct {
	rules   map[int]set
	updates [][]int
}

func load(inputPath string) (*input, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	input := input{
		rules:   map[int]set{},
		updates: [][]int{},
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		parts := strings.Split(text, "|")

		before, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse before value: %w", err)
		}

		after, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse after value: %w", err)
		}

		if _, ok := input.rules[before]; !ok {
			input.rules[before] = set{}
		}

		input.rules[before][after] = true
	}

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ",")

		update := make([]int, 0, len(parts))

		for _, part := range parts {
			page, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("failed to parse page value: %w", err)
			}

			update = append(update, page)
		}

		input.updates = append(input.updates, update)
	}

	return &input, nil
}

func (i *input) isValid(update []int) bool {
	for p, page := range update {
		// find any rules for pages after the current page
		pages, ok := i.rules[page]
		if !ok {
			continue
		}

		// check if any pages before the current page are invalid
		for _, before := range update[:p] {
			if pages[before] {
				return false
			}
		}
	}

	return true
}

func (i *input) part1() int {
	sum := 0

	for _, update := range i.updates {
		if i.isValid(update) {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func (i *input) part2() int {
	sum := 0

	for _, update := range i.updates {
		if !i.isValid(update) {
			slices.SortStableFunc(update, func(a, b int) int {
				pages, ok := i.rules[a]
				if !ok {
					return 0
				}

				if pages[b] {
					return -1
				}

				return 0
			})

			sum += update[len(update)/2]
		}
	}

	return sum
}
