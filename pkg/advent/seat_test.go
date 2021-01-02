package advent

import "testing"

func TestSeat(t *testing.T) {
	testCases := []struct {
		input  string
		row    int
		column int
		id     int
	}{
		{input: "FBFBBFFRLR", row: 44, column: 5, id: 357},
		{input: "BFFFBBFRRR", row: 70, column: 7, id: 567},
		{input: "FFFBBBFRRR", row: 14, column: 7, id: 119},
		{input: "BBFFBBFRLL", row: 102, column: 4, id: 820},
	}

	for _, tc := range testCases {
		s := NewSeat(tc.input)

		t.Run("row", func(t *testing.T) {
			if s.Row != tc.row {
				t.Fatalf("%v: got %#v, want %#v", tc.input, s.Row, tc.row)
			}
		})

		t.Run("column", func(t *testing.T) {
			if s.Column != tc.column {
				t.Fatalf("%v: got %#v, want %#v", tc.input, s.Column, tc.column)
			}
		})

		t.Run("id", func(t *testing.T) {
			if s.ID != tc.id {
				t.Fatalf("%v: got %#v, want %#v", tc.input, s.ID, tc.id)
			}
		})
	}
}
