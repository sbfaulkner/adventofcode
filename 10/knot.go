package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

const knotSize = 256

var finalLengths = []int{17, 31, 73, 47, 23}

type knot struct {
	list    []int
	lengths []int
	pos     int
	skip    int
}

func newKnot(f *os.File) *knot {
	k := knot{}

	k.list = make([]int, knotSize)

	for i := 0; i < knotSize; i++ {
		k.list[i] = i
	}

	s := bufio.NewScanner(f)
	s.Scan()

	string := s.Text()

	k.lengths = make([]int, len(string), len(string)+len(finalLengths))

	for i, ch := range string {
		k.lengths[i] = int(ch)
	}

	k.lengths = append(k.lengths, finalLengths...)

	return &k
}

func (k *knot) hash() string {
	for i := 0; i < 64; i++ {
		for _, l := range k.lengths {
			k.reverse(l)
			k.pos = (k.pos + l + k.skip) % knotSize
			k.skip++
		}
	}

	d := make([]byte, knotSize/16)

	for i := 0; i < knotSize; i++ {
		d[i/16] = d[i/16] ^ byte(k.list[i])
	}

	return hex.EncodeToString(d)
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

	fmt.Println(k.hash())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
