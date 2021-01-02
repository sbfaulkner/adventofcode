package advent

import (
	"strings"
	"testing"
)

func TestLoadProgram(t *testing.T) {
	in := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

	want := []Instruction{
		{Op: "nop", Arg: +0},
		{Op: "acc", Arg: +1},
		{Op: "jmp", Arg: +4},
		{Op: "acc", Arg: +3},
		{Op: "jmp", Arg: -3},
		{Op: "acc", Arg: -99},
		{Op: "acc", Arg: +1},
		{Op: "jmp", Arg: -4},
		{Op: "acc", Arg: +6},
	}

	got, err := LoadProgram(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	if len(*got) != len(want) {
		t.Fatalf("len: got %d, want %d", len(*got), len(want))
	}

	for ii, i := range *got {
		if i.Op != want[ii].Op {
			t.Errorf("Op[%d]: got %#v, want %#v", ii, i.Op, want[ii].Op)
		}

		if i.Arg != want[ii].Arg {
			t.Errorf("Arg[%d]: got %#v, want %#v", ii, i.Arg, want[ii].Arg)
		}
	}
}
