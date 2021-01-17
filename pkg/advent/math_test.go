package advent

import (
	"testing"
)

func TestMath(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, tc := range testCases {
		got, err := NewMathProblem(tc.input).Evaluate(LeftToRightPrecedence)
		if err != nil {
			t.Errorf("%s: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("%s: got %d, want %d", tc.input, got, tc.want)
		}
	}
}

func TestAdvancedMath(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, tc := range testCases {
		got, err := NewMathProblem(tc.input).Evaluate(AdvancedMathPrecedence)
		if err != nil {
			t.Errorf("%s: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("%s: got %d, want %d", tc.input, got, tc.want)
		}
	}
}
