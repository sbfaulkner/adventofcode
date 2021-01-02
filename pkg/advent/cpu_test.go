package advent

import (
	"strings"
	"testing"
)

func TestCPU(t *testing.T) {
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

	p, err := LoadProgram(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	cpu := CPU{}

	t.Run("failure", func(t *testing.T) {
		got := cpu.Execute(*p, cpu.DetectLoop())

		if got != false {
			t.Errorf("Execute: got %#v, want %#v", got, false)
		}

		if cpu.ACC != 5 {
			t.Errorf("ACC: got %d, want %d", cpu.ACC, 5)
		}
	})
}
