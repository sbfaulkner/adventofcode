package advent

import (
	"strings"
	"testing"
)

func TestReadDeclarations(t *testing.T) {
	in := `abc

a
b
c

ab
ac

a
a
a
a

b
`

	groups, err := ReadDeclarations(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		len, any, all int
	}{
		{1, 3, 3},
		{3, 3, 0},
		{2, 3, 1},
		{4, 1, 1},
		{1, 1, 1},
	}

	t.Run("groups", func(t *testing.T) {
		got := len(groups)

		if got != len(testCases) {
			t.Errorf("len: got %#v, want %#v", got, len(testCases))
		}
	})

	for i, tc := range testCases {
		g := groups[i]

		t.Run("len", func(t *testing.T) {
			got := len(g)

			if got != tc.len {
				t.Errorf("len[%d]: got %#v, want %#v", i, got, tc.len)
			}
		})

		t.Run("any", func(t *testing.T) {
			got := g.Any()

			if got != tc.any {
				t.Errorf("any[%d]: got %#v, want %#v", i, got, tc.any)
			}
		})

		t.Run("all", func(t *testing.T) {
			got := g.All()

			if got != tc.all {
				t.Errorf("all[%d]: got %#v, want %#v", i, got, tc.all)
			}
		})
	}
}
