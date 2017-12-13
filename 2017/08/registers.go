package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type cpu struct {
	registers map[string]int
}

func newCPU() *cpu {
	return &cpu{registers: map[string]int{}}
}

var lineRegexp = regexp.MustCompile(`^([a-z]+) (inc|dec) (-?[0-9]+) if ([a-z]+) (==|!=|>=|<=|<|>) (-?[0-9]+)$`)

func (cpu *cpu) execute(instruction string) {
	m := lineRegexp.FindStringSubmatch(instruction)

	if cpu.condition(m[4], m[5], m[6]) {
		cpu.perform(m[1], m[2], m[3])
	}
}

func (cpu *cpu) condition(register, operator, operand string) bool {
	r := cpu.registers[register]
	o, err := strconv.Atoi(operand)
	check(err)

	switch operator {
	case "==":
		return r == o
	case "!=":
		return r != o
	case "<":
		return r < o
	case "<=":
		return r <= o
	case ">":
		return r > o
	case ">=":
		return r >= o
	default:
		panic("invalid operator")
	}
}

func (cpu *cpu) perform(register, operator, operand string) {
	r := cpu.registers[register]
	o, err := strconv.Atoi(operand)
	check(err)

	switch operator {
	case "inc":
		r = r + o
	case "dec":
		r = r - o
	default:
		panic("invalid operator")
	}

	cpu.registers[register] = r
}

func (cpu *cpu) max() int {
	m := 0

	for _, v := range cpu.registers {
		if v > m {
			m = v
		}
	}

	return m
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	cpu := newCPU()
	s := bufio.NewScanner(f)

	for s.Scan() {
		instruction := s.Text()

		cpu.execute(instruction)
	}

	fmt.Println(cpu.registers)
	fmt.Println(cpu.max())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
