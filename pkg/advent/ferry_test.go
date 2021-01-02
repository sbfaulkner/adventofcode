package advent

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadFerry(t *testing.T) {
	in := `L.
.#
`

	got, err := ReadFerry(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	want := Ferry{
		[]rune("L."),
		[]rune(".#"),
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got %#v, want %#v", *got, want)
	}
}

func TestEvolve(t *testing.T) {
	ferry := &Ferry{
		[]rune("L.LL.LL.LL"),
		[]rune("LLLLLLL.LL"),
		[]rune("L.L.L..L.."),
		[]rune("LLLL.LL.LL"),
		[]rune("L.LL.LL.LL"),
		[]rune("L.LLLLL.LL"),
		[]rune("..L.L....."),
		[]rune("LLLLLLLLLL"),
		[]rune("L.LLLLLL.L"),
		[]rune("L.LLLLL.LL"),
	}

	want := []struct {
		ferry    Ferry
		evolved  bool
		occupied int
	}{
		{
			Ferry{
				[]rune("#.##.##.##"),
				[]rune("#######.##"),
				[]rune("#.#.#..#.."),
				[]rune("####.##.##"),
				[]rune("#.##.##.##"),
				[]rune("#.#####.##"),
				[]rune("..#.#....."),
				[]rune("##########"),
				[]rune("#.######.#"),
				[]rune("#.#####.##"),
			},
			true,
			71,
		},
		{
			Ferry{
				[]rune("#.LL.L#.##"),
				[]rune("#LLLLLL.L#"),
				[]rune("L.L.L..L.."),
				[]rune("#LLL.LL.L#"),
				[]rune("#.LL.LL.LL"),
				[]rune("#.LLLL#.##"),
				[]rune("..L.L....."),
				[]rune("#LLLLLLLL#"),
				[]rune("#.LLLLLL.L"),
				[]rune("#.#LLLL.##"),
			},
			true,
			20,
		},
		{
			Ferry{
				[]rune("#.##.L#.##"),
				[]rune("#L###LL.L#"),
				[]rune("L.#.#..#.."),
				[]rune("#L##.##.L#"),
				[]rune("#.##.LL.LL"),
				[]rune("#.###L#.##"),
				[]rune("..#.#....."),
				[]rune("#L######L#"),
				[]rune("#.LL###L.L"),
				[]rune("#.#L###.##"),
			},
			true,
			51,
		},
		{
			Ferry{
				[]rune("#.#L.L#.##"),
				[]rune("#LLL#LL.L#"),
				[]rune("L.L.L..#.."),
				[]rune("#LLL.##.L#"),
				[]rune("#.LL.LL.LL"),
				[]rune("#.LL#L#.##"),
				[]rune("..L.L....."),
				[]rune("#L#LLLL#L#"),
				[]rune("#.LLLLLL.L"),
				[]rune("#.#L#L#.##"),
			},
			true,
			30,
		},
		{
			Ferry{
				[]rune("#.#L.L#.##"),
				[]rune("#LLL#LL.L#"),
				[]rune("L.#.L..#.."),
				[]rune("#L##.##.L#"),
				[]rune("#.#L.LL.LL"),
				[]rune("#.#L#L#.##"),
				[]rune("..L.L....."),
				[]rune("#L#L##L#L#"),
				[]rune("#.LLLLLL.L"),
				[]rune("#.#L#L#.##"),
			},
			true,
			37,
		},
		{
			Ferry{
				[]rune("#.#L.L#.##"),
				[]rune("#LLL#LL.L#"),
				[]rune("L.#.L..#.."),
				[]rune("#L##.##.L#"),
				[]rune("#.#L.LL.LL"),
				[]rune("#.#L#L#.##"),
				[]rune("..L.L....."),
				[]rune("#L#L##L#L#"),
				[]rune("#.LLLLLL.L"),
				[]rune("#.#L#L#.##"),
			},
			false,
			37,
		},
	}

	for g, f := range want {
		var occupied int
		var evolved bool

		ferry, occupied, evolved = ferry.Evolve(1, 4)

		if !reflect.DeepEqual(*ferry, f.ferry) {
			t.Errorf("Generation %d: got %s, want %s", g+2, *ferry, f.ferry)
		}

		if evolved != f.evolved {
			t.Errorf("Generation %d: evolved got %#v, want %#v", g+2, evolved, f.evolved)
		}

		if occupied != f.occupied {
			t.Errorf("Generation %d: occupied got %#v, want %#v", g+2, occupied, f.occupied)
		}
	}
}

func TestEnhancedEvolve(t *testing.T) {
	ferry := &Ferry{
		[]rune("L.LL.LL.LL"),
		[]rune("LLLLLLL.LL"),
		[]rune("L.L.L..L.."),
		[]rune("LLLL.LL.LL"),
		[]rune("L.LL.LL.LL"),
		[]rune("L.LLLLL.LL"),
		[]rune("..L.L....."),
		[]rune("LLLLLLLLLL"),
		[]rune("L.LLLLLL.L"),
		[]rune("L.LLLLL.LL"),
	}

	want := []struct {
		ferry    Ferry
		evolved  bool
		occupied int
	}{
		{
			Ferry{
				[]rune("#.##.##.##"),
				[]rune("#######.##"),
				[]rune("#.#.#..#.."),
				[]rune("####.##.##"),
				[]rune("#.##.##.##"),
				[]rune("#.#####.##"),
				[]rune("..#.#....."),
				[]rune("##########"),
				[]rune("#.######.#"),
				[]rune("#.#####.##"),
			},
			true,
			71,
		},
		{
			Ferry{
				[]rune("#.LL.LL.L#"),
				[]rune("#LLLLLL.LL"),
				[]rune("L.L.L..L.."),
				[]rune("LLLL.LL.LL"),
				[]rune("L.LL.LL.LL"),
				[]rune("L.LLLLL.LL"),
				[]rune("..L.L....."),
				[]rune("LLLLLLLLL#"),
				[]rune("#.LLLLLL.L"),
				[]rune("#.LLLLL.L#"),
			},
			true,
			7,
		},
		{
			Ferry{
				[]rune("#.L#.##.L#"),
				[]rune("#L#####.LL"),
				[]rune("L.#.#..#.."),
				[]rune("##L#.##.##"),
				[]rune("#.##.#L.##"),
				[]rune("#.#####.#L"),
				[]rune("..#.#....."),
				[]rune("LLL####LL#"),
				[]rune("#.L#####.L"),
				[]rune("#.L####.L#"),
			},
			true,
			53,
		},
		{
			Ferry{
				[]rune("#.L#.L#.L#"),
				[]rune("#LLLLLL.LL"),
				[]rune("L.L.L..#.."),
				[]rune("##LL.LL.L#"),
				[]rune("L.LL.LL.L#"),
				[]rune("#.LLLLL.LL"),
				[]rune("..L.L....."),
				[]rune("LLLLLLLLL#"),
				[]rune("#.LLLLL#.L"),
				[]rune("#.L#LL#.L#"),
			},
			true,
			18,
		},
		{
			Ferry{
				[]rune("#.L#.L#.L#"),
				[]rune("#LLLLLL.LL"),
				[]rune("L.L.L..#.."),
				[]rune("##L#.#L.L#"),
				[]rune("L.L#.#L.L#"),
				[]rune("#.L####.LL"),
				[]rune("..#.#....."),
				[]rune("LLL###LLL#"),
				[]rune("#.LLLLL#.L"),
				[]rune("#.L#LL#.L#"),
			},
			true,
			31,
		},
		{
			Ferry{
				[]rune("#.L#.L#.L#"),
				[]rune("#LLLLLL.LL"),
				[]rune("L.L.L..#.."),
				[]rune("##L#.#L.L#"),
				[]rune("L.L#.LL.L#"),
				[]rune("#.LLLL#.LL"),
				[]rune("..#.L....."),
				[]rune("LLL###LLL#"),
				[]rune("#.LLLLL#.L"),
				[]rune("#.L#LL#.L#"),
			},
			true,
			26,
		},
		{
			Ferry{
				[]rune("#.L#.L#.L#"),
				[]rune("#LLLLLL.LL"),
				[]rune("L.L.L..#.."),
				[]rune("##L#.#L.L#"),
				[]rune("L.L#.LL.L#"),
				[]rune("#.LLLL#.LL"),
				[]rune("..#.L....."),
				[]rune("LLL###LLL#"),
				[]rune("#.LLLLL#.L"),
				[]rune("#.L#LL#.L#"),
			},
			false,
			26,
		},
	}

	for g, f := range want {
		var occupied int
		var evolved bool

		ferry, occupied, evolved = ferry.Evolve(9, 5)

		if !reflect.DeepEqual(ferry, f.ferry) {
			t.Errorf("Generation %d: got %s, want %s", g+2, ferry, f.ferry)
		}

		if evolved != f.evolved {
			t.Errorf("Generation %d: evolved got %#v, want %#v", g+2, evolved, f.evolved)
		}

		if occupied != f.occupied {
			t.Errorf("Generation %d: occupied got %#v, want %#v", g+2, occupied, f.occupied)
		}
	}
}
