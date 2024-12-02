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
	safety  []bool
}

func load(inputPath string) (*input, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}

	input := input{
		reports: [][]int{},
		safety:  []bool{},
	}

	scanner := bufio.NewScanner(file)

	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		fields := strings.Fields(line)

		input.reports = append(input.reports, make([]int, 0, len(fields)))
		input.safety = append(input.safety, true)

		previousDiff := 0

		for f, field := range fields {
			level, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("failed to parse field value %s: %w", field, err)
			}

			input.reports[r] = append(input.reports[r], level)

			// we're already iterating over the fields, so pre-calculate the diffs/safety
			if f > 0 {
				d := input.reports[r][f] - input.reports[r][f-1]

				// if the diff is too big or zero, the report is not safe
				if d < -3 || d > 3 || d == 0 {
					input.safety[r] = false
					break
				}

				// if the diff is the opposite sign of the previous diff, the report is not safe
				if f > 1 && d*previousDiff < 0 {
					input.safety[r] = false
					break
				}

				previousDiff = d
			}

			input.safety[r] = true
		}
	}

	return &input, nil
}

func (i *input) safe() int {
	safe := 0

	for _, s := range i.safety {
		if s {
			safe++
		}
	}

	return safe
}
