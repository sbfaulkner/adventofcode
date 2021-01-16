package advent

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

const (
	eofToken = iota
	intToken
	opToken
	lparenToken
	rparenToken
)

type token struct {
	t int
	i int
	s string
}

// MathProblem is an expression to be solved
type MathProblem struct {
	rd    *strings.Reader
	err   error
	depth int
	tok   token
}

// ReadMathHomework reads a series of math problems
func ReadMathHomework(rd io.Reader) ([]*MathProblem, error) {
	h := []*MathProblem{}

	s := bufio.NewScanner(rd)

	for s.Scan() {
		h = append(h, NewMathProblem(s.Text()))
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return h, nil
}

// NewMathProblem creates a new math problem
func NewMathProblem(problem string) *MathProblem {
	p := MathProblem{
		rd: strings.NewReader(problem),
	}
	return &p
}

func (p *MathProblem) getToken() bool {
	for {
		ch, _, err := p.rd.ReadRune()
		if err == io.EOF {
			p.tok.t = eofToken
			return true
		}

		if err != nil {
			p.err = err
			return false
		}

		if unicode.IsSpace(ch) {
			continue
		}

		p.tok.s = string(ch)

		switch {
		case unicode.IsDigit(ch):
			p.tok.t = intToken
			return p.getNumber()

		case strings.ContainsRune("+*", ch):
			p.tok.t = opToken
			return true

		case ch == '(':
			p.tok.t = lparenToken
			return true

		case ch == ')':
			p.tok.t = rparenToken
			return true

		default:
			p.err = fmt.Errorf("syntax error at %c", ch)
			return false
		}
	}
}

func (p *MathProblem) getNumber() bool {
	for {
		ch, _, err := p.rd.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			p.err = err
			return false
		}

		if !unicode.IsDigit(ch) {
			err = p.rd.UnreadRune()
			if err != nil {
				p.err = err
				return false
			}
			break
		}

		p.tok.s += string(ch)
	}

	i, err := strconv.Atoi(p.tok.s)
	if err != nil {
		p.err = err
		return false
	}

	p.tok.i = i

	return true
}

// OperatorPrecedence to use when evaluating an expression
type OperatorPrecedence func(o1, o2 string) bool

// LeftToRightPrecedence defines left-to-right operator precedence
func LeftToRightPrecedence(o1, o2 string) bool {
	return true
}

// AdvancedMathPrecedence defines plus-over-times operator precedence
func AdvancedMathPrecedence(o1, o2 string) bool {
	if o1 == o2 {
		return true
	}

	return o1 == "+"
}

// Evaluate solves a math problem
func (p *MathProblem) Evaluate(f OperatorPrecedence) (int, error) {
	return p.evaluate(f)
}

func (p *MathProblem) evaluate(f OperatorPrecedence) (int, error) {
	p.depth++
	defer func() { p.depth-- }()

	for p.getToken() {
		switch p.tok.t {
		case intToken:
			return p.evaluateLeft(f, p.tok.i)

		case lparenToken:
			l, err := p.evaluate(f)
			if err != nil {
				return 0, err
			}
			return p.evaluateLeft(f, l)

		default:
			return 0, fmt.Errorf("syntax error at %s (expected integer)", p.tok.s)
		}
	}

	return 0, p.err
}

func (p *MathProblem) evaluateLeft(f OperatorPrecedence, l int) (int, error) {
	if !p.getToken() {
		return l, p.err
	}

	if p.depth > 1 {
		if p.tok.t == rparenToken {
			return l, nil
		}
	} else {
		if p.tok.t == eofToken {
			return l, nil
		}
	}

	if p.tok.t != opToken {
		return 0, fmt.Errorf("syntax error at %s (expected operator)", p.tok.s)
	}

	s, err := p.evaluateOp(f, l, p.tok.s)
	if err != nil {
		return s, err
	}

	return p.evaluateLeft(f, s)
}

func (p *MathProblem) evaluateOp(f OperatorPrecedence, l int, op string) (int, error) {
	if !p.getToken() {
		return 0, p.err
	}

	switch p.tok.t {
	case intToken:
		return p.evaluateExpression(f, l, op, p.tok.i)

	case lparenToken:
		r, err := p.evaluate(f)
		if err != nil {
			return 0, err
		}
		return p.evaluateExpression(f, l, op, r)
	}

	return 0, fmt.Errorf("syntax error at %s (expected rval)", p.tok.s)
}

func (p *MathProblem) evaluateExpression(f OperatorPrecedence, l int, op string, r int) (int, error) {
	switch op {
	case "+":
		return l + r, nil
	case "*":
		return l * r, nil
	}

	return 0, fmt.Errorf("unknown operator at %s", op)
}
