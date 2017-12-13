package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type program struct {
	name     string
	weight   int
	total    int
	parent   *program
	children []*program
	t        *tower
}

type tower map[string]*program

func (t *tower) programByName(name string) *program {
	p, ok := (*t)[name]

	if !ok {
		p = t.newProgram(name)
	}

	return p
}

func (t tower) findBottom() *program {
	var p *program

	for _, p = range t {
		break
	}

	for p.parent != nil {
		p = p.parent
	}

	return p
}

func (t *tower) newProgram(name string) *program {
	p := &program{name: name, t: t, children: []*program{}}

	(*t)[name] = p

	return p
}

func (p *program) addChildren(names []string) {
	for _, n := range names {
		h := p.t.programByName(n)
		h.parent = p
		p.children = append(p.children, h)
	}
}

func (p *program) updateTotal() int {
	p.total = p.weight

	for _, c := range p.children {
		p.total = p.total + c.updateTotal()
	}

	return p.total
}

func (p *program) findImbalance() *program {
	weights := map[int][]*program{}

	for _, c := range p.children {
		if i := c.findImbalance(); i != nil {
			return i
		}

		weights[c.total] = append(weights[c.total], c)
	}

	for _, s := range weights {
		if len(s) == 1 && len(weights) > 1 {
			return s[0]
		}
	}

	return nil
}

func (p *program) requiredWeight() int {
	for _, c := range p.parent.children {
		if c.total != p.total {
			return p.weight + c.total - p.total
		}
	}

	return p.weight
}

var lineRegexp = regexp.MustCompile(`^([a-z]+) \(([0-9]+)\)(?: -> ([a-z]+(?:, [a-z]+)*))?$`)

func main() {
	f, err := os.Open("input.txt")
	check(err)

	t := tower{}
	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()

		matches := lineRegexp.FindStringSubmatch(line)

		name := matches[1]
		weight, err := strconv.Atoi(matches[2])
		check(err)
		names := []string{}
		if matches[3] != "" {
			names = strings.Split(matches[3], ", ")
		}

		p := t.programByName(name)
		p.weight = weight
		p.addChildren(names)
	}

	bottom := t.findBottom()
	bottom.updateTotal()
	fmt.Println("bottom is", bottom)

	incorrect := bottom.findImbalance()
	fmt.Println("incorrect is", incorrect)

	weight := incorrect.requiredWeight()
	fmt.Println("required weight is", weight)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
