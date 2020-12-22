package passport

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Validate is a password field validation function
type Validate func(Passport) bool

// Requirements is a collection of field validations
type Requirements []Validate

// RequirePresent determines if the specified field is present
func RequirePresent(f string) Validate {
	return func(p Passport) bool {
		_, ok := p[f]
		return ok
	}
}

// Range defines a min and max value
type Range struct{ Min, Max int }

// RequireIntegerInRange determines if the specified field has a value within a range
func RequireIntegerInRange(f string, r Range) Validate {
	return func(p Passport) bool {
		v, err := strconv.Atoi(p[f])
		if err != nil {
			return false
		}

		return r.Min <= v && v <= r.Max
	}
}

// RequireMeasurementInRange determines if the specified field has a measurement within a set of ranges
func RequireMeasurementInRange(f string, ranges map[string]Range) Validate {
	return func(p Passport) bool {
		var v int
		var u string

		n, err := fmt.Sscanf(p[f], "%d%s", &v, &u)
		if err != nil || n != 2 {
			return false
		}

		r := ranges[u]

		return r.Min <= v && v <= r.Max
	}
}

var rgbRegexp = regexp.MustCompile(`\A#[0-9a-f]{6}\z`)

// RequireMatch determines if the specified field matches the provided regular expression
func RequireMatch(f string, str string) Validate {
	re := regexp.MustCompile(str)
	return func(p Passport) bool {
		return re.MatchString(p[f])
	}
}

// RequireOneOf determins if the specified field has one of the provided values
func RequireOneOf(f string, values ...string) Validate {
	return func(p Passport) bool {
		for _, v := range values {
			if p[f] == v {
				return true
			}
		}

		return false
	}
}

// RequireFields is a simple set of required fields
var RequireFields = Requirements{
	RequirePresent("byr"),
	RequirePresent("iyr"),
	RequirePresent("eyr"),
	RequirePresent("hgt"),
	RequirePresent("hcl"),
	RequirePresent("ecl"),
	RequirePresent("pid"),
}

// RequireValidFields is a more complex set of field validations
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
var RequireValidFields = Requirements{
	RequireIntegerInRange("byr", Range{Min: 1920, Max: 2002}),
	RequireIntegerInRange("iyr", Range{Min: 2010, Max: 2020}),
	RequireIntegerInRange("eyr", Range{Min: 2020, Max: 2030}),
	RequireMeasurementInRange("hgt", map[string]Range{"cm": {Min: 150, Max: 193}, "in": {Min: 59, Max: 76}}),
	RequireMatch("hcl", `\A#[0-9a-f]{6}\z`),
	RequireOneOf("ecl", "amb", "blu", "brn", "gry", "grn", "hzl", "oth"),
	RequireMatch("pid", `\A[0-9]{9}\z`),
}

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
		passports = append(passports, p)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return passports, nil
}

// ValidPassports counts the valid passports based on the provided requirements
func ValidPassports(passports []*Passport, r Requirements) int {
	v := 0

	for _, p := range passports {
		if p.Valid(r) {
			v++
		}
	}

	return v
}

// Valid determines if the passport is valid
func (p Passport) Valid(r Requirements) bool {
	for _, v := range r {
		if !v(p) {
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
