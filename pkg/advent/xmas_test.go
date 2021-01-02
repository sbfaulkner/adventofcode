package advent

import (
	"strings"
	"testing"
)

func TestXMAS(t *testing.T) {
	in := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

	xmas, err := LoadXMAS(5, strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("FindInvalid", func(t *testing.T) {
		got := xmas.FindInvalid()

		if got != 127 {
			t.Errorf("got %d, want %d", got, 127)
		}
	})

	t.Run("FindRangeMinMax", func(t *testing.T) {
		gotMin, gotMax := xmas.FindRangeMinMax(127)

		wantMin := 15
		wantMax := 47

		if gotMin != wantMin {
			t.Errorf("got %#v, want %#v", gotMin, wantMin)
		}

		if gotMax != wantMax {
			t.Errorf("got %#v, want %#v", gotMax, wantMax)
		}
	})
}
