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
			ts:       939,
			schedule: Schedule{7, 13, 0, 0, 59, 0, 31, 19},
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
}
