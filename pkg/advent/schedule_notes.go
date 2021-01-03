package advent

import (
	"fmt"
	"io"
)

// ScheduleNotes for shuttles
type ScheduleNotes struct {
	ts int
	s  schedule
}

// Read the shuttle schedule notes
func (n *ScheduleNotes) Read(rd io.Reader) error {
	if _, err := fmt.Fscanln(rd, &n.ts); err != nil {
		return err
	}

	return n.s.read(rd)
}

// FindBus finds the first bus after the timestamp according to the schedule
func (n ScheduleNotes) FindBus() (int, int) {
	var id, wait int

	for _, i := range n.s.ids {
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
func (n ScheduleNotes) FindSyncTime() uint64 {
	ts := uint64(n.s.ids[0])

	lcm := uint64(1)

	n.s.each(func(i, id int) {
		for (ts+uint64(i))%uint64(id) != 0 {
			ts += lcm
		}
		lcm *= uint64(id)
	})

	return ts
}
