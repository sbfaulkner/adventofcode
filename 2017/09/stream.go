package main

import (
	"bufio"
	"fmt"
	"os"
)

type streamProcessor func(*stream, rune) streamProcessor

type stream struct {
	scanner   *bufio.Scanner
	processor streamProcessor
	group     int
	score     int
}

func newStream(f *os.File) *stream {
	s := stream{processor: findGroup}

	s.scanner = bufio.NewScanner(f)
	s.scanner.Split(bufio.ScanRunes)

	return &s
}

func findGroup(s *stream, ch rune) streamProcessor {
	switch ch {
	case '{':
		s.group++
	case '}':
		s.score = s.score + s.group
		s.group--
	case '<':
		return collectGarbage
	}

	return findGroup
}

func collectGarbage(s *stream, ch rune) streamProcessor {
	switch ch {
	case '!':
		return skipGarbage
	case '>':
		return findGroup
	}

	return collectGarbage
}

func skipGarbage(s *stream, ch rune) streamProcessor {
	return collectGarbage
}

func (s *stream) calculateScore() int {
	for s.scanner.Scan() {
		ch := s.scanner.Text()
		s.processor = s.processor(s, []rune(ch)[0])
	}

	return s.score
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	s := newStream(f)
	score := s.calculateScore()

	fmt.Println(score)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
