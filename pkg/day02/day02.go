package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type input struct {
	reports [][]int
}

func load(inputPath string) (*input, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	input := input{
		reports: [][]int{},
	}

	scanner := bufio.NewScanner(file)

	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		fields := strings.Fields(line)

		input.reports = append(input.reports, make([]int, 0, len(fields)))

		for _, field := range fields {
			level, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("failed to parse field value %s: %w", field, err)
			}

			input.reports[r] = append(input.reports[r], level)
		}
	}

	return &input, nil
}

func safe(report []int) bool {
	previousDiff := 0

	for f := range report {
		if f > 0 {
			d := report[f] - report[f-1]

			// if the diff is too big or zero, the report is not safe
			if d < -3 || d > 3 || d == 0 {
				return false
			}

			// if the diff is the opposite sign of the previous diff, the report is not safe
			if f > 1 && d*previousDiff < 0 {
				return false
			}

			previousDiff = d
		}
	}

	return true
}

func (i *input) safe() int {
	count := 0

	for _, report := range i.reports {
		if safe(report) {
			count++
		}
	}

	return count
}

func (i *input) safeWithDampener() int {
	count := 0

	for _, report := range i.reports {
		if safe(report) {
			count++
			continue
		}

		// if the report is not safe, we need to check if it's safe with one level removed
		dampened := make([]int, len(report)-1)

		for f := range report {
			copy(dampened, report[:f])
			copy(dampened[f:], report[f+1:])

			if safe(dampened) {
				count++
				break
			}
		}
	}

	return count
}
