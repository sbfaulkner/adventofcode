package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	count := 0

	for s.Scan() {
		line := s.Text()

		words := strings.Split(line, " ")
		m := map[string]bool{}

		for _, word := range words {
			m[word] = true
		}

		if len(m) == len(words) {
			count++
		}
	}

	fmt.Println(count)
}
