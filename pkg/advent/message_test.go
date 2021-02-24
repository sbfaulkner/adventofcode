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
		t.Errorf("CountValid: got %d, want %d", got, 2)
	}
}

func TestMessagesCountValid(t *testing.T) {
	in := `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

	rd := strings.NewReader(in)

	r, m := ReadMessages(rd)

	got := r.CountValid(m)
	if got != 3 {
		t.Errorf("original: got %d, want %d", got, 3)
	}

	r[8].alternatives = [][]int{{42}, {42, 8}}
	r[11].alternatives = [][]int{{42, 31}, {42, 11, 31}}

	got = r.CountValid(m)
	if got != 12 {
		t.Errorf("modified: got %d, want %d", got, 12)
	}
}
