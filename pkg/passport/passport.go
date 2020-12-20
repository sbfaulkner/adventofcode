package passport

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// type field struct {
// 	name     string
// 	required bool
// }

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

// // byr (Birth Year) - four digits; at least 1920 and at most 2002.
// // iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// // eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// // hgt (Height) - a number followed by either cm or in:
// // If cm, the number must be at least 150 and at most 193.
// // If in, the number must be at least 59 and at most 76.
// // hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// // pid (Passport ID) - a nine-digit number, including leading zeroes.
// // cid (Country ID) - ignored, missing or not.

// var validFields = []field{
// 	{name: "byr", required: true},
// 	{name: "iyr", required: true},
// 	{name: "eyr", required: true},
// 	{name: "hgt", required: true},
// 	{name: "hcl", required: true},
// 	{name: "ecl", required: true},
// 	{name: "pid", required: true},
// 	{name: "cid"},
// }

// func (p passport) validate(fields []field) (bool, []error) {
// 	valid := true
// 	errors := []error{}

// 	for _, f := range fields {
// 		if f.required {
// 			_, ok := p[f.name]
// 			if !ok {
// 				valid = false
// 				errors = append(errors, fmt.Errorf("%s is required", f.name))
// 			}
// 		}
// 	}

// 	return valid, errors
// }

// Passport holds sets of field-value pairs
type Passport map[string]string

// NewPassport creates a new Passport
func NewPassport(data string) *Passport {
	p := Passport{}

	fields := strings.Fields(data)

	for _, f := range fields {
		s := strings.Split(f, ":")
		p[s[0]] = s[1]
	}

	return &p
}

// ReadPassports reads all passports from a Reader
func ReadPassports(rd io.Reader) ([]*Passport, error) {
	passports := []*Passport{}

	s := NewScanner(rd)

	for s.Scan() {
		p := NewPassport(s.Text())
		if p.Valid() {
			passports = append(passports, p)
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return passports, nil
}

// Valid determines if the passport is valid
func (p Passport) Valid() bool {
	for _, f := range requiredFields {
		if _, bOk := p[f]; !bOk {
			return false
		}
	}
	return true
}

// scanPassport is a split function for a Scanner that returns a full set of passport lines
func scanPassport(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		// We have a full double-newline-terminated passport.
		return i + 2, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		l := len(data)
		if l > 0 && data[l-1] == '\n' {
			return l, data[:l-1], nil
		}
		return l, data, nil
	}
	// Request more data.
	return 0, nil, nil
}

// NewScanner creates a Passport Scanner
func NewScanner(rd io.Reader) *bufio.Scanner {
	s := bufio.NewScanner(rd)
	s.Split(scanPassport)
	return s
}
