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
	st    []token
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

func (p *MathProblem) pushToken() {
	p.st = append(p.st, p.tok)
}

func (p *MathProblem) pushInt(i int) {
	p.tok.t = intToken
	p.tok.i = i
	p.pushToken()
}

func (p *MathProblem) popToken() token {
	t := p.st[len(p.st)-1]
	p.st = p.st[0 : len(p.st)-1]
	return t
}

func (p *MathProblem) popInt() int {
	return p.popToken().i
}

func (p *MathProblem) popString() string {
	return p.popToken().s
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
	if err := p.evaluate(f); err != nil {
		return 0, err
	}
	return p.popInt(), nil
}

func (p *MathProblem) evaluate(f OperatorPrecedence) error {
	p.depth++
	defer func() { p.depth-- }()

	for p.getToken() {
		switch p.tok.t {
		case intToken:
			p.pushToken()
			return p.evaluateLeft(f)

		case lparenToken:
			err := p.evaluate(f)
			if err != nil {
				return err
			}
			return p.evaluateLeft(f)

		default:
			return fmt.Errorf("syntax error at %s (expected integer)", p.tok.s)
		}
	}

	return p.err
}

func (p *MathProblem) evaluateLeft(f OperatorPrecedence) error {
	if !p.getToken() {
		return p.err
	}

	if p.depth > 1 {
		if p.tok.t == rparenToken {
			return nil
		}
	} else {
		if p.tok.t == eofToken {
			return nil
		}
	}

	if p.tok.t != opToken {
		return fmt.Errorf("syntax error at %s (expected operator)", p.tok.s)
	}

	p.pushToken()
	if err := p.evaluateOp(f); err != nil {
		return err
	}

	return p.evaluateLeft(f)
}

func (p *MathProblem) evaluateOp(f OperatorPrecedence) error {
	if !p.getToken() {
		return p.err
	}

	switch p.tok.t {
	case intToken:
		p.pushToken()
		return p.evaluateExpression(f)

	case lparenToken:
		err := p.evaluate(f)
		if err != nil {
			return err
		}
		return p.evaluateExpression(f)
	}

	return fmt.Errorf("syntax error at %s (expected rval)", p.tok.s)
}

func (p *MathProblem) evaluateExpression(f OperatorPrecedence) error {
	l := p.popInt()
	op := p.popString()
	r := p.popInt()

	switch op {
	case "+":
		p.pushInt(l + r)
	case "*":
		p.pushInt(l * r)
	default:
		return fmt.Errorf("unknown operator at %s", op)
	}

	return nil
}
