package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func sledPolicy(in string) bool {
	var atLeast, atMost int
	var ch rune
	var p string

	_, err := fmt.Sscanf(in, "%d-%d %c: %s", &atLeast, &atMost, &ch, &p)
	if err != nil {
		log.Fatal(err)
	}

	c := 0

	for _, r := range p {
		if r == ch {
			c = c + 1
		}
	}

	return c >= atLeast && c <= atMost
}

func tobogganPolicy(in string) bool {
	var i1, i2 int
	var ch rune
	var p string

	_, err := fmt.Sscanf(in, "%d-%d %c: %s", &i1, &i2, &ch, &p)
	if err != nil {
		log.Fatal(err)
	}

	r := []rune(p)

	return r[i1-1] != r[i2-1] && (r[i1-1] == ch || r[i2-1] == ch)
}

// CountValid returns the number of valid passwords
func CountValid(rd io.Reader, p func(string) bool) int {
	count := 0

	scanner := bufio.NewScanner(rd)

	for scanner.Scan() {
		if p(scanner.Text()) {
			count = count + 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return count
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

	s := CountValid(file, sledPolicy)

	log.Println(s)

	file.Seek(0, io.SeekStart)

	t := CountValid(file, tobogganPolicy)

	log.Println(t)
}
