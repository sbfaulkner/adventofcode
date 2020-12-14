package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type slope struct{ dx, dy int }

// TreeCount returns the number of trees on the trajectories
func TreeCount(rd io.Reader, slopes []slope) []int {
	trees := [][]rune{}

	scanner := bufio.NewScanner(rd)

	for scanner.Scan() {
		trees = append(trees, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	counts := make([]int, len(slopes))

	for i, s := range slopes {
		var x, y int

		for y < len(trees) {
			if trees[y][x] == '#' {
				counts[i] = counts[i] + 1
			}

			y = y + s.dy
			x = (x + s.dx) % len(trees[0])
		}
	}

	return counts
}

// TreeCountProduct returns the product of the tree counts for the trajectories
func TreeCountProduct(rd io.Reader, slopes []slope) int {
	p := 1

	for _, c := range TreeCount(rd, slopes) {
		p = p * c
	}

	return p
}

func main() {
	_, sourcePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Unable to get source path")
	}

	sourceDir := filepath.Dir(sourcePath)

	file, err := os.Open(path.Join(sourceDir, "input"))
	if err != nil {
		log.Fatal(err)
	}

	t := TreeCount(file, []slope{{3, 1}})

	log.Println(t[0])

	file.Seek(0, io.SeekStart)

	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	p := TreeCountProduct(file, slopes)
	log.Println(p)
}
