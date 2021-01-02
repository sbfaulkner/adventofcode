package advent

// CPU for handheld game console
type CPU struct {
	ACC int
	pc  int
}

// Execute the provided program
func (cpu *CPU) Execute(p Program, debug Debugger) bool {
	for cpu.pc < len(p) {
		if !debug() {
			return false
		}
		cpu.execute(p[cpu.pc])
	}

	return true
}

func (cpu *CPU) execute(i Instruction) {
	switch i.Op {
	case "acc":
		cpu.ACC += i.Arg
		cpu.pc++
	case "jmp":
		cpu.pc += i.Arg
	case "nop":
		cpu.pc++
	}
}

// DetectLoop function to stop at first instruction already executed
func (cpu *CPU) DetectLoop() Debugger {
	breakpoints := make(map[int]bool)

	return func() bool {
		if breakpoints[cpu.pc] {
			return false
		}
		breakpoints[cpu.pc] = true
		return true
	}
}
