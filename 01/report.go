package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
)

// Fix returns the product of the two entries that add to 2020
func Fix(rd io.Reader) int {
	values := readIntegers(rd)

	for _, i1 := range values {
		for _, i2 := range values {
			if i1+i2 == 2020 {
				return i1 * i2
			}
		}
	}

	return -1
}

// Fix3 returns the product of the three entries that add to 2020
func Fix3(rd io.Reader) int {
	values := readIntegers(rd)

	for _, i1 := range values {
		for _, i2 := range values {
			for _, i3 := range values {
				if i1+i2+i3 == 2020 {
					return i1 * i2 * i3
				}
			}
		}
	}

	return -1
}

func readIntegers(rd io.Reader) []int {
	values := []int{}

	scanner := bufio.NewScanner(rd)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values
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

	fix := Fix(file)

	log.Println(fix)

	file.Seek(0, io.SeekStart)

	fix3 := Fix3(file)

	log.Println(fix3)
}
