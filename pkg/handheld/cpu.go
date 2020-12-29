package handheld

import "fmt"

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

// Trace currently executing program
func (cpu *CPU) Trace(p Program, debug Debugger) Debugger {
	return func() bool {
		i := p[cpu.pc]
		fmt.Printf("%04d: %s %+d [acc: %+d]\n", cpu.pc, i.Op, i.Arg, cpu.ACC)
		if !debug() {
			fmt.Println("BREAK")
			return false
		}
		return true
	}
}
