package advent

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// Program is a set of instructions
type Program []Instruction

// LoadProgram loads a program from the specified reader
func LoadProgram(rd io.Reader) (*Program, error) {
	p := Program{}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		err := p.appendInstruction(s.Text())
		if err != nil {
			return nil, err
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Program) appendInstruction(i string) error {
	f := strings.Fields(i)

	arg, err := strconv.Atoi(f[1])
	if err != nil {
		return err
	}

	*p = append(*p, Instruction{Op: f[0], Arg: arg})

	return nil
}
