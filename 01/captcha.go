package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	s := bufio.NewScanner(f)
	s.Scan()
	t := s.Text()

	sum := 0

	for i, r := range t {
		j := (i + len(t)/2) % len(t)
		if byte(r) == []byte(t)[j] {
			sum = sum + int(r-'0')
		}
	}

	fmt.Println(sum)
}
