package advent

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	id       int
	ch       rune
	subrules [][]int
}

// MessageRules is a collection of message validation rules
type MessageRules map[int]*rule

var ruleExpression = regexp.MustCompile(`\A(?P<id>\d+): (?:"(?P<ch>.)"|(?P<rules>.+))\z`)
var ruleIDIndex = ruleExpression.SubexpIndex("id")
var ruleChIndex = ruleExpression.SubexpIndex("ch")
var ruleRulesIndex = ruleExpression.SubexpIndex("rules")

func parseRule(s string) *rule {
	m := ruleExpression.FindStringSubmatch(s)

	id, _ := strconv.Atoi(m[ruleIDIndex]) // ignore error
	ch := []rune(m[ruleChIndex])
	rules := m[ruleRulesIndex]

	rule := rule{id: id}

	if len(ch) > 0 {
		rule.ch = ch[0]
	} else {
		alternatives := strings.Split(rules, " | ")
		rule.subrules = make([][]int, 0, len(alternatives))
		for ai, a := range alternatives {
			subrules := strings.Split(a, " ")
			rule.subrules = append(rule.subrules, make([]int, 0, len(subrules)))
			for _, sr := range subrules {
				sri, _ := strconv.Atoi(sr) // ignore error
				rule.subrules[ai] = append(rule.subrules[ai], sri)
			}
		}
	}

	return &rule
}

func (r *MessageRules) scan(s *bufio.Scanner) {
	for s.Scan() {
		t := s.Text()
		if len(t) == 0 {
			break
		}

		p := parseRule(t)
		(*r)[p.id] = p
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func (r MessageRules) matchRule(id int, m string) int {
	rule := r[id]
	if rule.ch != 0 {
		if rule.ch == []rune(m)[0] {
			return 1
		}
		return 0
	}

Alternatives:
	for _, alt := range rule.subrules {
		count := 0

		for _, sub := range alt {
			c := r.matchRule(sub, m[count:])
			if c == 0 {
				continue Alternatives
			}

			count += c
		}

		return count
	}

	return 0
}

func (r MessageRules) valid(m string) bool {
	return r.matchRule(0, m) == len(m)
}

// CountValid returns the number of valid messages
func (r MessageRules) CountValid(messages []string) int {
	count := 0
	for _, m := range messages {
		if r.valid(m) {
			count++
		}
	}
	return count
}

// ReadMessages reads the rules and the messages
func ReadMessages(rd io.Reader) (MessageRules, []string) {
	s := bufio.NewScanner(rd)

	r := MessageRules{}

	r.scan(s)

	m := []string{}

	for s.Scan() {
		m = append(m, s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return r, m
}
