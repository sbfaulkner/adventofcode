package advent

import (
	"strings"
	"testing"
)

func TestInitialize(t *testing.T) {
	dp := DockingProgram{}

	in := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`

	if err := dp.Initialize(strings.NewReader(in)); err != nil {
		t.Fatal(err)
	}

	got := dp.CheckSum()

	if got != 165 {
		t.Errorf("CheckSum: got %d, want %d", got, 165)
	}
}
