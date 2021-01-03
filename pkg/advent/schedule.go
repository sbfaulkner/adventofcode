package advent

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ReadSchedule reads the scheduled shuttles
func ReadSchedule(rd io.Reader) (map[int]int, error) {
	var ts int
	var s string

	if _, err := fmt.Fscanln(rd, &ts); err != nil {
		return nil, err
	}

	if _, err := fmt.Fscanln(rd, &s); err != nil && err != io.EOF {
		return nil, err
	}

	schedule := map[int]int{}

	for _, f := range strings.Split(s, ",") {
		if f == "x" {
			continue
		}

		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}

		schedule[n] = nextMultiple(ts, n) - ts
	}

	return schedule, nil
}

func nextMultiple(i int, m int) int {
	r := i % m

	if r == 0 {
		return i
	}

	return i + m - r
}
