package advent

import (
	"bufio"
	"io"
)

// Declaration is a collection of answers for the customs declaration form
type Declaration map[rune]bool

// Group is a collection of Declarations for a group o passengers
type Group []*Declaration

// NewDeclaration creates a customs declaration form
func NewDeclaration(in string) *Declaration {
	d := Declaration{}

	for _, q := range in {
		d[q] = true
	}

	return &d
}

// ReadDeclarations reads the declarations from the specified reader
func ReadDeclarations(rd io.Reader) ([]Group, error) {
	groups := []Group{}
	group := make(Group, 0)

	s := bufio.NewScanner(rd)

	for s.Scan() {
		t := s.Text()

		if len(t) > 0 {
			group = append(group, NewDeclaration(t))
		} else {
			groups = append(groups, group)
			group = make(Group, 0)
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	if len(group) > 0 {
		groups = append(groups, group)
	}

	return groups, nil
}

// Any is the number of questions which any member of the group answered yes to
func (g Group) Any() int {
	answers := map[rune]bool{}

	for _, d := range g {
		for q := range *d {
			answers[q] = true
		}
	}

	return len(answers)
}

// All is the number of questions which all members of the group answered yes to
func (g Group) All() int {
	answers := map[rune]int{}

	for _, d := range g {
		for q := range *d {
			answers[q]++
		}
	}

	c := 0

	for _, a := range answers {
		if a == len(g) {
			c++
		}
	}

	return c
}
