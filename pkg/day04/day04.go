package day04

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type input struct {
	rows []string
}

func load(inputPath string) (*input, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	input := input{
		rows: []string{},
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input.rows = append(input.rows, scanner.Text())
	}

	return &input, nil
}

func (i *input) part1() int {
	count := 0

	for r, row := range i.rows {
		for c := range row {
			count += i.search(r, c, "XMAS")
		}
	}

	return count
}

type direction struct {
	dr int
	dc int
}

var DIRECTIONS = []direction{
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
}

func (i *input) search(r, c int, target string) int {
	if i.rows[r][c] != target[0] {
		return 0
	}

	count := 0

	for _, dir := range DIRECTIONS {
		if i.matchDirection(r+dir.dr, c+dir.dc, target[1:], dir) {
			count++
		}
	}

	return count
}

func (i *input) matchDirection(r, c int, target string, dir direction) bool {
	if r < 0 || r >= len(i.rows) || c < 0 || c >= len(i.rows[r]) {
		return false
	}

	if i.rows[r][c] != target[0] {
		return false
	}

	if len(target) == 1 {
		return true
	}

	return i.matchDirection(r+dir.dr, c+dir.dc, target[1:], dir)
}

func (i *input) part2() int {
	count := 0

	for r, row := range i.rows[1 : len(i.rows)-1] {
		for c := range row[1 : len(row)-1] {
			if i.xMasAt(r+1, c+1) {
				count++
			}
		}
	}

	return count
}

var SEQUENCES = [][]byte{
	{'M', 'M', 'S', 'S'},
	{'S', 'M', 'M', 'S'},
	{'S', 'S', 'M', 'M'},
	{'M', 'S', 'S', 'M'},
}

func (i *input) xMasAt(r, c int) bool {
	if i.rows[r][c] != 'A' {
		return false
	}

	sequence := []byte{i.rows[r-1][c-1], i.rows[r-1][c+1], i.rows[r+1][c+1], i.rows[r+1][c-1]}

	for _, seq := range SEQUENCES {
		if bytes.Equal(sequence, seq) {
			return true
		}
	}

	return false
}
