package advent

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// TicketNotes taken for high-speed train
type TicketNotes struct {
	rules     TicketFieldRules
	ticket    Ticket
	nearby    []Ticket
	errorRate int
}

// TicketFieldRules maps a ticket field class to a set of ranges
type TicketFieldRules map[string]TicketFieldRule

// TicketFieldRule is a set of valid ranges for a ticket field
type TicketFieldRule []Range

// Ticket is a collection field values for a ticket
type Ticket []int

// ReadTicketNotes reads ticket notes from the provided reader
func ReadTicketNotes(rd io.Reader) (*TicketNotes, error) {
	notes := TicketNotes{}

	s := bufio.NewScanner(rd)

	if err := notes.readRules(s); err != nil {
		return nil, err
	}

	if err := expectText(s, "your ticket:"); err != nil {
		return nil, err
	}

	t, err := readTicket(s)
	if err != nil {
		return nil, err
	}

	notes.ticket = *t

	if err := expectText(s, "", "nearby tickets:"); err != nil {
		return nil, err
	}

	for {
		t, err := readTicket(s)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		e := 0

		for _, f := range *t {
			v := notes.rules.findPossibleFields(f)
			if len(v) == 0 {
				e += f
			}
		}

		if e == 0 {
			notes.nearby = append(notes.nearby, *t)
		} else {
			notes.errorRate += e
		}
	}

	return &notes, nil
}

var ticketFieldRuleRegexp = regexp.MustCompile("^(.+): ([0-9]+-[0-9]+) or ([0-9]+-[0-9]+)$")

func (n *TicketNotes) readRules(s *bufio.Scanner) error {
	n.rules = make(TicketFieldRules)

	for s.Scan() {
		t := s.Text()
		if t == "" {
			break
		}

		m := ticketFieldRuleRegexp.FindStringSubmatch(t)
		if m == nil {
			return fmt.Errorf("invalid rule – %s", t)
		}

		if err := n.rules.addRange(m[1], m[2]); err != nil {
			return err
		}

		if err := n.rules.addRange(m[1], m[3]); err != nil {
			return err
		}
	}

	return s.Err()
}

// TicketScanningErrorRate returns to total of the field values in error
func (n TicketNotes) TicketScanningErrorRate() int {
	return n.errorRate
}

type stringset []string

func (ss *stringset) add(s string) {
	*ss = append(*ss, s)
}

func (ss *stringset) remove(s string) {
	for i, str := range *ss {
		if str == s {
			*ss = append((*ss)[0:i], (*ss)[i+1:]...)
			break
		}
	}
}

// GetDepartureFieldCheckSum returns the product of the departure fields
func (n TicketNotes) GetDepartureFieldCheckSum() int {
	product := 1

	for i, f := range n.GetTicketFieldOrder() {
		if strings.HasPrefix(f, "departure ") {
			product *= n.ticket[i]
		}
	}

	return product
}

// GetTicketFieldOrder determines the order of the fields on the tickets
func (n TicketNotes) GetTicketFieldOrder() []string {
	choices := make([]stringset, len(n.ticket))

	for i, f := range n.ticket {
		choices[i] = stringset{}

	RULES:
		for c, r := range n.rules {
			if !r.allows(f) {
				continue
			}

			for _, t := range n.nearby {
				if !r.allows(t[i]) {
					continue RULES
				}
			}

			choices[i].add(c)
		}
	}

	fields := make([]string, len(n.ticket))

	for {
		count := 0

		for i, c := range choices {
			if len(c) > 1 {
				continue
			}

			fields[i] = c[0]
			count++

			for ii := range choices {
				if ii == i {
					continue
				}

				choices[ii].remove(c[0])
			}
		}

		if count == len(n.ticket) {
			break
		}
	}

	return fields
}

func expectText(s *bufio.Scanner, expected ...string) error {
	for _, e := range expected {
		if !s.Scan() {
			if err := s.Err(); err != nil {
				return err
			}
			return io.EOF
		}

		t := s.Text()
		if t != e {
			return fmt.Errorf("Syntax error at %q (expected %q)", t, e)
		}
	}

	return nil
}

func readTicket(s *bufio.Scanner) (*Ticket, error) {
	if !s.Scan() {
		if err := s.Err(); err != nil {
			return nil, err
		}
		return nil, io.EOF
	}

	t := Ticket{}

	for _, f := range strings.Split(s.Text(), ",") {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		t = append(t, n)
	}

	return &t, nil
}

func (r *TicketFieldRules) addRange(c string, rs string) error {
	s := strings.Split(rs, "-")

	min, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}

	max, err := strconv.Atoi(s[1])
	if err != nil {
		return err
	}

	(*r)[c] = append((*r)[c], Range{Min: min, Max: max})

	return nil
}

func (r TicketFieldRules) findPossibleFields(v int) []string {
	matches := []string{}

	for c, rule := range r {
		if rule.allows(v) {
			matches = append(matches, c)
		}
	}

	return matches
}

func (r TicketFieldRule) allows(v int) bool {
	for _, ra := range r {
		if ra.Includes(v) {
			return true
		}
	}
	return false
}
