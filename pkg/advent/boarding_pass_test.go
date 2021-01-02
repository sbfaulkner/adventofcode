package advent

import "testing"

func TestBoardingPass(t *testing.T) {
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
		bp := NewBoardingPass(tc.input)

		t.Run("row", func(t *testing.T) {
			if bp.Row != tc.row {
				t.Fatalf("%v: got %#v, want %#v", tc.input, bp.Row, tc.row)
			}
		})

		t.Run("column", func(t *testing.T) {
			if bp.Column != tc.column {
				t.Fatalf("%v: got %#v, want %#v", tc.input, bp.Column, tc.column)
			}
		})

		t.Run("id", func(t *testing.T) {
			if bp.ID != tc.id {
				t.Fatalf("%v: got %#v, want %#v", tc.input, bp.ID, tc.id)
			}
		})
	}
}
