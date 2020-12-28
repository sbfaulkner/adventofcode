package bag

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

// Rule sets expectations for the content of a bag of a given color
type Rule map[string]int

// Rules is a full set of content rules
type Rules map[string]*Rule

// ReadRules loads a set of rules from a reader
func ReadRules(rd io.Reader) (*Rules, error) {
	rules := Rules{}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		c, r, err := ParseRule(s.Text())
		if err != nil {
			return nil, err
		}
		rules[c] = r
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return &rules, nil
}

var ruleRegexp = regexp.MustCompile(`\A([^ ]+ [^ ]+) bags contain (.*)\.\z`)
var contentsRegexp = regexp.MustCompile(`([0-9]+) ([^ ]+ [^ ]+) bags?`)

// ParseRule parses a color and a content rulde from line of input
func ParseRule(rule string) (string, *Rule, error) {
	r := Rule{}

	matches := ruleRegexp.FindStringSubmatch(rule)
	for _, m := range contentsRegexp.FindAllStringSubmatch(matches[2], -1) {
		n, err := strconv.Atoi(m[1])
		if err != nil {
			return "", nil, err
		}
		r[m[2]] = n
	}

	return matches[1], &r, nil
}

// FindContaining returns colors that can (eventually) contain another
func (r Rules)FindContaining(color string) []string {
	containing := []string{}

	for c, rule := range r {
		if rule.contains(color, &r) {
			containing = append(containing, c)
		}
	}

	return containing
}

func (r Rule) contains(color string, rules* Rules) bool {
	_, bOK := r[color]
	if bOK {
		return true
	}

	for c := range r {
		if (*rules)[c].contains(color, rules) {
			return true
		}
	}

	return false
}