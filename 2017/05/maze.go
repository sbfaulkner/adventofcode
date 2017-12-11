package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	maze := []int{}

	for s.Scan() {
		line := s.Text()

		i, err := strconv.Atoi(line)
		check(err)

		maze = append(maze, i)
	}

	c := 0
	j := 0

	for o := 0; o >= 0 && o < len(maze); o = o + j {
		j = maze[o]
		if j > 2 {
			maze[o]--
		} else {
			maze[o]++
		}
		c++
	}

	fmt.Println(c)
}
