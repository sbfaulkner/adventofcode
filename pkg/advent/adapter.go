package advent

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

// Adapter represented as a joltage-rating
type Adapter struct {
	joltage int
	chains int
}

// Adapters is a chain of adapters
type Adapters []*Adapter

// ReadAdapters from reader
func ReadAdapters(rd io.Reader) (Adapters, error) {
	adapters := Adapters{}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		j, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}

		adapters = append(adapters, &Adapter{joltage: j})
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	sort.Slice(adapters, func(i, j int) bool { return adapters[i].joltage < adapters[j].joltage })

	return adapters, nil
}

// CountAdapters counts the number of adapters with each joltage-difference
func (a Adapters) CountAdapters() map[int]int {
	c := make(map[int]int, len(a))

	j := 0

	for _, adapter := range a {
		d := adapter.joltage - j
		c[d] = c[d] + 1
		j = adapter.joltage
	}

	c[3] = c[3] + 1

	return c
}

// CountChains counts the number of chains possible for the set of adapters
func (a *Adapter) CountChains(adapters Adapters) int {
	if a.chains > 0 {
		return a.chains
	}

	for i, adapter := range adapters {
		if adapter.joltage > a.joltage + 3 {
			break
		}

		a.chains += adapter.CountChains(adapters[i+1:])
	}

	if a.chains == 0 {
		a.chains++
	}

	return a.chains
}
