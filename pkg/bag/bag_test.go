package bag

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadRules(t *testing.T) {
	in := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

	want := Rules{
		"light red": &Rule{"bright white": 1, "muted yellow": 2},
		"dark orange": &Rule{"bright white": 3, "muted yellow": 4},
		"bright white": &Rule{"shiny gold": 1},
		"muted yellow": &Rule{"shiny gold": 2, "faded blue": 9},
		"shiny gold": &Rule{"dark olive": 1, "vibrant plum": 2},
		"dark olive": &Rule{"faded blue": 3, "dotted black": 4},
		"vibrant plum": &Rule{"faded blue": 5, "dotted black": 6},
		"faded blue": &Rule{},
		"dotted black": &Rule{},
	}

	rules, err := ReadRules(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("len", func(t *testing.T) {
		if len(*rules) != len(want) {
			t.Fatalf("got %#v, want %#v", len(*rules), len(want))
		}
	})

	for c, r := range want {
		t.Run(c, func(t *testing.T) {
			g := (*rules)[c]
			if !reflect.DeepEqual(*g, *r) {
				t.Fatalf("got %#v, want %#v", *g, *r)
			}
		})
	}

	t.Run("FindContaining", func(t *testing.T) {
		got := rules.FindContaining("shiny gold")
		if len(got) != 4 {
			t.Fatalf("len: got %d, want %d", len(got), 4)
		}
	})
}

func TestParseRule(t *testing.T) {
	testCases := []struct {
		in    string
		color string
		rule  Rule
	}{
		{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"light red",
			Rule{"bright white": 1, "muted yellow": 2},
		},
		{
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"dark orange",
			Rule{"bright white": 3, "muted yellow": 4},
		},
		{
			"bright white bags contain 1 shiny gold bag.",
			"bright white",
			Rule{"shiny gold": 1},
		},
		{
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"muted yellow",
			Rule{"shiny gold": 2, "faded blue": 9},
		},
		{
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"shiny gold",
			Rule{"dark olive": 1, "vibrant plum": 2},
		},
		{
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"dark olive",
			Rule{"faded blue": 3, "dotted black": 4},
		},
		{
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"vibrant plum",
			Rule{"faded blue": 5, "dotted black": 6},
		},
		{
			"faded blue bags contain no other bags.",
			"faded blue",
			Rule{},
		},
		{
			"dotted black bags contain no other bags.",
			"dotted black",
			Rule{},
		},
	}

	for _, tc := range testCases {
		c, r, err := ParseRule(tc.in)
		if err != nil {
			t.Error(err)
			continue
		}

		t.Run("color", func(t *testing.T) {
			if c != tc.color {
				t.Errorf("got %#v, want %#v", c, tc.color)
			}
		})

		t.Run(c, func(t *testing.T) {
			if len(*r) != len(tc.rule) {
				t.Fatalf("len: got %d, want %d", len(*r), len(tc.rule))
			}

			for c, want := range tc.rule {
				if (*r)[c] != want {
					t.Errorf("%s: got %d, want %d", c, (*r)[c], want)
				}
			}
		})
	}
}

func TestContentsOf(t *testing.T) {
	in := `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`

	rules, err := ReadRules(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	contents := rules.ContentsOf("shiny gold")
	if len(contents) != 126 {
		t.Errorf("len: got %d, expected %d", len(contents), 126)
	}
}