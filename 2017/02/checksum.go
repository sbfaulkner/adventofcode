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

const maxuint = ^uint(0)
const maxint = int(maxuint >> 1)

func main() {
	f, err := os.Open("input.txt")
	check(err)

	sum := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()

		values := strings.Split(line, "\t")

		row := []int{}

		for _, value := range values {
			v, err := strconv.Atoi(value)
			check(err)

			row = append(row, v)
		}

		for i, ii := range row {
			for j, jj := range row {
				if i == j {
					continue
				}

				if ii%jj == 0 {
					sum = sum + ii/jj
					break
				}
			}
		}
	}

	fmt.Println(sum)
}
