package advent

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

// Policy is a password validation function
type Policy func(string) bool

// SledPolicy validates passwords using the sled rental company policy
func SledPolicy(in string) bool {
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

// TobogganPolicy validates passwords using the North Pole Toboggan Rental Shop policy
func TobogganPolicy(in string) bool {
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

// Database is a collection of password policies and passwords
type Database []string

// ReadDatabase reads a password policy database
func ReadDatabase(rd io.Reader) (*Database, error) {
	d := make(Database, 0)

	scanner := bufio.NewScanner(rd)

	for scanner.Scan() {
		d = append(d, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &d, nil
}

// CountValid returns the number of valid passwords
func (db Database) CountValid(p Policy) int {
	count := 0

	for _, e := range db {
		if p(e) {
			count++
		}
	}

	return count
}
