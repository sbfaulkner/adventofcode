package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type memory []int

func newMemory(b []string) memory {
	m := memory(make([]int, len(b)))

	for i, v := range b {
		n, err := strconv.Atoi(v)
		check(err)
		m[i] = n
	}

	return m
}

func (m memory) fingerprint() string {
	return fmt.Sprint([]int(m))
}

func (m *memory) balance() {
	i := m.maxBankIndex()

	b := (*m)[i]
	(*m)[i] = 0

	for b > 0 {
		i = (i + 1) % len(*m)
		(*m)[i]++
		b--
	}
}

func (m memory) maxBankIndex() int {
	max := 0

	for i := range m {
		if m[i] > m[max] {
			max = i
		}
	}

	return max
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	s := bufio.NewScanner(f)
	s.Scan()

	b := strings.Split(s.Text(), "\t")
	m := newMemory(b)

	states := map[string]bool{}

	for i := 0; ; i++ {
		if _, ok := states[m.fingerprint()]; ok {
			fmt.Println(i)
			break
		}

		states[m.fingerprint()] = true
		m.balance()
	}

	states = map[string]bool{}

	for i := 0; ; i++ {
		if _, ok := states[m.fingerprint()]; ok {
			fmt.Println(i)
			break
		}

		states[m.fingerprint()] = true
		m.balance()
	}
}
