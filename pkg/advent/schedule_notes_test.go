package advent

import (
	"reflect"
	"strings"
	"testing"
)

func TestScheduleNotes(t *testing.T) {
	in := `939
7,13,x,x,59,x,31,19
`

	n := ScheduleNotes{}

	if err := n.Read(strings.NewReader(in)); err != nil {
		t.Fatal(err)
	}

	t.Run("Read", func(t *testing.T) {
		want := ScheduleNotes{
			ts: 939,
			s: schedule{
				index: []int{0, 1, 4, 6, 7},
				ids:   map[int]int{0: 7, 1: 13, 4: 59, 6: 31, 7: 19},
			},
		}

		if !reflect.DeepEqual(n, want) {
			t.Errorf("got %#v, want %#v", n, want)
		}
	})

	t.Run("FindBus", func(t *testing.T) {
		id, wait := n.FindBus()

		if id != 59 {
			t.Errorf("id: got %#v, want %#v", id, 59)
		}

		if wait != 5 {
			t.Errorf("wait: got %#v, want %#v", wait, 5)
		}
	})

	t.Run("FindSyncTime", func(t *testing.T) {
		got := n.FindSyncTime()

		if got != 1068781 {
			t.Errorf("got %#v, want %#v", got, uint64(1068781))
		}
	})
}

func TestFindSyncTime(t *testing.T) {
	testCases := []struct {
		n    ScheduleNotes
		want uint64
	}{
		{
			ScheduleNotes{s: schedule{index: []int{0, 2, 3}, ids: map[int]int{0: 17, 2: 13, 3: 19}}},
			3417,
		},
		{
			ScheduleNotes{s: schedule{index: []int{0, 1, 2, 3}, ids: map[int]int{0: 67, 1: 7, 2: 59, 3: 61}}},
			754018,
		},
		{
			ScheduleNotes{s: schedule{index: []int{0, 2, 3, 4}, ids: map[int]int{0: 67, 2: 7, 3: 59, 4: 61}}},
			779210,
		},
		{
			ScheduleNotes{s: schedule{index: []int{0, 1, 3, 4}, ids: map[int]int{0: 67, 1: 7, 3: 59, 4: 61}}},
			1261476,
		},
		{
			ScheduleNotes{s: schedule{index: []int{0, 1, 2, 3}, ids: map[int]int{0: 1789, 1: 37, 2: 47, 3: 1889}}},
			1202161486,
		},
	}

	for _, tc := range testCases {
		got := tc.n.FindSyncTime()
		if got != tc.want {
			t.Errorf("got %#v, want %#v", got, tc.want)
		}
	}
}
