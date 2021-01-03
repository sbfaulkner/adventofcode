package advent

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ScheduleNotes for shuttles
type ScheduleNotes struct {
	ts       int
	schedule Schedule
}

// Schedule for shuttles
type Schedule []int

// Read the shuttle schedule notes
func (n *ScheduleNotes) Read(rd io.Reader) error {
	if _, err := fmt.Fscanln(rd, &n.ts); err != nil {
		return err
	}

	var s string

	if _, err := fmt.Fscanln(rd, &s); err != nil && err != io.EOF {
		return err
	}

	buses := strings.Split(s, ",")

	n.schedule = make(Schedule, len(buses))

	for i, f := range buses {
		if f == "x" {
			continue
		}

		id, err := strconv.Atoi(f)
		if err != nil {
			return err
		}

		n.schedule[i] = id
	}

	return nil
}

// FindBus finds the first bus after the timestamp according to the schedule
func (n ScheduleNotes) FindBus() (int, int) {
	var id, wait int

	for _, i := range n.schedule {
		if i == 0 {
			continue
		}

		w := nextMultiple(n.ts, i) - n.ts

		if id == 0 || w < wait {
			id = i
			wait = w
		}
	}

	return id, wait
}

func nextMultiple(i int, m int) int {
	r := i % m

	if r == 0 {
		return i
	}

	return i + m - r
}
