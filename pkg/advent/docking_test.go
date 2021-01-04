package advent

import (
	"strings"
	"testing"
)

func TestInitializeV1(t *testing.T) {
	dp := DockingProgram{}

	in := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`

	if err := dp.Initialize(strings.NewReader(in), &DecoderV1{}); err != nil {
		t.Fatal(err)
	}

	got := dp.CheckSum()

	if got != 165 {
		t.Errorf("CheckSum: got %d, want %d", got, 165)
	}
}

func TestInitializeV2(t *testing.T) {
	dp := DockingProgram{}

	in := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`

	if err := dp.Initialize(strings.NewReader(in), &DecoderV2{}); err != nil {
		t.Fatal(err)
	}

	got := dp.CheckSum()

	if got != 208 {
		t.Errorf("CheckSum: got %d, want %d", got, 208)
	}
}
