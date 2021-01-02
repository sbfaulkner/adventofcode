package advent

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadSeating(t *testing.T) {
	in := `L.
.#
`

	got, err := ReadSeating(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	want := Seating{
		[]rune("L."),
		[]rune(".#"),
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got %#v, want %#v", *got, want)
	}
}

func TestEvolve(t *testing.T) {
	seating := &Seating{
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
		seating  Seating
		evolved  bool
		occupied int
	}{
		{
			Seating{
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
			Seating{
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
			Seating{
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
			Seating{
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
			Seating{
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
			Seating{
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

	for g, w := range want {
		var occupied int
		var evolved bool

		seating, occupied, evolved = seating.Evolve(1, 4)

		if !reflect.DeepEqual(*seating, w.seating) {
			t.Errorf("Generation %d: got %#v, want %#v", g+2, *seating, w.seating)
		}

		if evolved != w.evolved {
			t.Errorf("Generation %d: evolved got %#v, want %#v", g+2, evolved, w.evolved)
		}

		if occupied != w.occupied {
			t.Errorf("Generation %d: occupied got %#v, want %#v", g+2, occupied, w.occupied)
		}
	}
}

func TestEnhancedEvolve(t *testing.T) {
	seating := &Seating{
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
		seating  Seating
		evolved  bool
		occupied int
	}{
		{
			Seating{
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
			Seating{
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
			Seating{
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
			Seating{
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
			Seating{
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
			Seating{
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
			Seating{
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

	for g, w := range want {
		var occupied int
		var evolved bool

		seating, occupied, evolved = seating.Evolve(9, 5)

		if !reflect.DeepEqual(*seating, w.seating) {
			t.Errorf("Generation %d: got %#v, want %#v", g+2, *seating, w.seating)
		}

		if evolved != w.evolved {
			t.Errorf("Generation %d: evolved got %#v, want %#v", g+2, evolved, w.evolved)
		}

		if occupied != w.occupied {
			t.Errorf("Generation %d: occupied got %#v, want %#v", g+2, occupied, w.occupied)
		}
	}
}
