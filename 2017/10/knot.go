package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const knotSize = 256

type knot struct {
	list    [knotSize]int
	lengths []int
	pos     int
	skip    int
}

func newKnot(f *os.File) *knot {
	k := knot{}

	for i := 0; i < 256; i++ {
		k.list[i] = i
	}

	s := bufio.NewScanner(f)
	s.Scan()

	strings := strings.Split(s.Text(), ",")
	k.lengths = make([]int, len(strings))

	for i, str := range strings {
		l, err := strconv.Atoi(str)
		check(err)
		k.lengths[i] = l
	}

	return &k
}

func (k *knot) tie() {
	for _, l := range k.lengths {
		k.reverse(l)
		k.pos = k.pos + l + k.skip
		k.skip++
	}
}

func (k *knot) reverse(l int) {
	for i := 0; i < l/2; i++ {
		x := (k.pos + i) % knotSize
		y := (k.pos + l - 1 - i) % knotSize

		k.list[x], k.list[y] = k.list[y], k.list[x]
	}
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	k := newKnot(f)

	k.tie()

	fmt.Println(k.list[0] * k.list[1])
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
