package advent

import (
	"strings"
	"testing"
)

func TestMessages(t *testing.T) {
	in := `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

	rd := strings.NewReader(in)

	r, m := ReadMessages(rd)

	if len(r) != 6 {
		t.Errorf("len(rules): got %d, want %d", len(r), 6)
	}

	if len(m) != 5 {
		t.Errorf("len: got %d, want %d", len(m), 5)
	}

	want := []bool{true, false, true, false, false}

	for i, w := range want {
		got := r.valid(m[i])
		if got != w {
			t.Errorf("valid(%#v): got %#v, want %#v", m[i], got, want)
		}
	}

	got := r.CountValid(m)
	if got != 2 {
		t.Errorf("CountVaid: got %d, want %d", got, 2)
	}
}
