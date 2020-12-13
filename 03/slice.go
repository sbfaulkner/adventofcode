package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

const fabricSize = 1000
const inputName = "input"

func sourceFile() string {
	_, file, _, _ := runtime.Caller(0)
	return file
}

func sourceDir() string {
	return path.Dir(sourceFile())
}

func inputFile() string {
	return path.Join(sourceDir(), inputName)
}

type claim struct {
	id      int
	left    int
	top     int
	width   int
	height  int
	overlap bool
}

type fabric struct {
	pattern [][][]int
	claims  map[int]*claim
}

func (f *fabric) claim(text string) {
	c := claim{}

	fmt.Sscanf(text, "#%d @ %d,%d: %dx%d", &c.id, &c.left, &c.top, &c.width, &c.height)

	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			f.pattern[c.top+y][c.left+x] = append(f.pattern[c.top+y][c.left+x], c.id)

			if len(f.pattern[c.top+y][c.left+x]) > 1 {
				previous := f.pattern[c.top+y][c.left+x][len(f.pattern[c.top+y][c.left+x])-2]

				c.overlap = true
				f.claims[previous].overlap = true
			}
		}
	}

	f.claims[c.id] = &c
}

func (f fabric) totalOverlap() int {
	total := 0

	for _, slice := range f.pattern {
		for _, claims := range slice {
			if len(claims) > 1 {
				total = total + 1
			}
		}
	}

	return total
}

func newFabric(size int) *fabric {
	f := fabric{
		pattern: make([][][]int, 0, size),
		claims:  map[int]*claim{},
	}

	for y := 0; y < size; y++ {
		f.pattern = append(f.pattern, make([][]int, size))
	}

	return &f
}

func main() {
	input, err := os.Open(inputFile())
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	fabric := newFabric(fabricSize)

	for scanner.Scan() {
		fabric.claim(scanner.Text())
	}

	fmt.Println("overlap", fabric.totalOverlap())

	for _, c := range fabric.claims {
		if !c.overlap {
			fmt.Println(c.id, "is isolated")
		}
	}
}
