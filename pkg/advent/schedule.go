package advent

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// ScheduleNotes for shuttles
type ScheduleNotes struct {
	ts       int
	schedule Schedule
}

// Schedule for shuttles
type Schedule struct {
	index []int
	ids   map[int]int
}

// Read the shuttle schedule notes
func (n *ScheduleNotes) Read(rd io.Reader) error {
	if _, err := fmt.Fscanln(rd, &n.ts); err != nil {
		return err
	}

	return n.schedule.read(rd)
}

func (s *Schedule) read(rd io.Reader) error {
	var schedule string

	if _, err := fmt.Fscanln(rd, &schedule); err != nil && err != io.EOF {
		return err
	}

	buses := strings.Split(schedule, ",")

	s.index = make([]int, 0, len(buses))
	s.ids = make(map[int]int, len(buses))

	for i, f := range buses {
		if f == "x" {
			continue
		}

		id, err := strconv.Atoi(f)
		if err != nil {
			return err
		}

		s.index = append(s.index, i)
		s.ids[i] = id
	}

	sort.Slice(s.index, func(i, j int) bool { return s.index[i] < s.index[j] })

	return nil
}

// FindBus finds the first bus after the timestamp according to the schedule
func (n ScheduleNotes) FindBus() (int, int) {
	var id, wait int

	for _, i := range n.schedule.ids {
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

// FindSyncTime finds the time the schedule is in sync with the route order
func (n ScheduleNotes) FindSyncTime() int64 {
	return n.schedule.findSyncTime()
}

func (s Schedule) findSyncTime() int64 {
	var ts int64
	time := int64(s.ids[s.index[0]])

TIME:
	for ; ; ts += time {
		for _, i := range s.index[1:] {
			id := s.ids[i]
			if (ts+int64(i))%int64(id) != 0 {
				continue TIME
			}
		}
		break
	}

	return ts
}
