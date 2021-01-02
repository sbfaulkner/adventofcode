package advent

import (
	"strings"
	"testing"
)

func TestSledPolicy(t *testing.T) {
	testCases := []struct {
		in   string
		want bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", true},
	}

	for _, tc := range testCases {
		got := SledPolicy(tc.in)
		if got != tc.want {
			t.Errorf("SledPolicy(%#v): got %#v, want %#v", tc.in, got, tc.want)
		}
	}
}

func TestTobogganPolicy(t *testing.T) {
	testCases := []struct {
		in   string
		want bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", false},
	}

	for _, tc := range testCases {
		got := TobogganPolicy(tc.in)
		if got != tc.want {
			t.Errorf("TobogganPolicy(%#v): got %#v, want %#v", tc.in, got, tc.want)
		}
	}
}

func TestCountValid(t *testing.T) {
	in := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`
	db, err := ReadDatabase(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name   string
		policy Policy
		want   int
	}{
		{"none", func(string) bool { return false }, 0},
		{"all", func(string) bool { return true }, 3},
		{"SledPolicy", SledPolicy, 2},
		{"TobogganPolicy", TobogganPolicy, 1},
	}

	for _, tc := range testCases {
		got := db.CountValid(tc.policy)
		if got != tc.want {
			t.Errorf("db.CountValid(%#v): got %#v, want %#v", tc.name, got, tc.want)
		}
	}
}
