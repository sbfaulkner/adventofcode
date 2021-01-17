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
type MathProblem string

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
	p := MathProblem(problem)
	return &p
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
	e := mathEvaluator{
		rd: strings.NewReader(string(*p)),
		f:  f,
	}

	if err := evaluateProblem(&e); err != nil {
		return 0, err
	}
	return e.popInt(), nil
}

type mathEvaluator struct {
	rd    *strings.Reader
	err   error
	depth int
	tok   token
	st    []token
	f     OperatorPrecedence
}

func (e *mathEvaluator) getToken() bool {
	for {
		ch, _, err := e.rd.ReadRune()
		if err == io.EOF {
			e.tok.t = eofToken
			return true
		}

		if err != nil {
			e.err = err
			return false
		}

		if unicode.IsSpace(ch) {
			continue
		}

		e.tok.s = string(ch)

		switch {
		case unicode.IsDigit(ch):
			e.tok.t = intToken
			return e.getNumber()

		case strings.ContainsRune("+*", ch):
			e.tok.t = opToken
			return true

		case ch == '(':
			e.tok.t = lparenToken
			return true

		case ch == ')':
			e.tok.t = rparenToken
			return true

		default:
			e.err = fmt.Errorf("syntax error at %c", ch)
			return false
		}
	}
}

func (e *mathEvaluator) getNumber() bool {
	for {
		ch, _, err := e.rd.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			e.err = err
			return false
		}

		if !unicode.IsDigit(ch) {
			err = e.rd.UnreadRune()
			if err != nil {
				e.err = err
				return false
			}
			break
		}

		e.tok.s += string(ch)
	}

	i, err := strconv.Atoi(e.tok.s)
	if err != nil {
		e.err = err
		return false
	}

	e.tok.i = i

	return true
}

func (e *mathEvaluator) pushToken() {
	e.st = append(e.st, e.tok)
}

func (e *mathEvaluator) pushInt(i int) {
	e.tok.t = intToken
	e.tok.i = i
	e.pushToken()
}

func (e *mathEvaluator) popToken() token {
	t := e.st[len(e.st)-1]
	e.st = e.st[0 : len(e.st)-1]
	return t
}

func (e *mathEvaluator) popInt() int {
	return e.popToken().i
}

func (e *mathEvaluator) popString() string {
	return e.popToken().s
}

func evaluateProblem(e *mathEvaluator) error {
	e.depth++
	defer func() { e.depth-- }()

	for e.getToken() {
		switch e.tok.t {
		case intToken:
			e.pushToken()
			return evaluateLeft(e)

		case lparenToken:
			err := evaluateProblem(e)
			if err != nil {
				return err
			}
			return evaluateLeft(e)

		default:
			return fmt.Errorf("syntax error at %s (expected integer)", e.tok.s)
		}
	}

	return e.err
}

func evaluateLeft(e *mathEvaluator) error {
	if !e.getToken() {
		return e.err
	}

	if e.depth > 1 {
		if e.tok.t == rparenToken {
			return nil
		}
	} else {
		if e.tok.t == eofToken {
			return nil
		}
	}

	if e.tok.t != opToken {
		return fmt.Errorf("syntax error at %s (expected operator)", e.tok.s)
	}

	e.pushToken()
	if err := evaluateOp(e); err != nil {
		return err
	}

	return evaluateLeft(e)
}

func evaluateOp(e *mathEvaluator) error {
	if !e.getToken() {
		return e.err
	}

	switch e.tok.t {
	case intToken:
		e.pushToken()
		return evaluateExpression(e)

	case lparenToken:
		err := evaluateProblem(e)
		if err != nil {
			return err
		}
		return evaluateExpression(e)
	}

	return fmt.Errorf("syntax error at %s (expected rval)", e.tok.s)
}

func evaluateExpression(e *mathEvaluator) error {
	l := e.popInt()
	op := e.popString()
	r := e.popInt()

	switch op {
	case "+":
		e.pushInt(l + r)
	case "*":
		e.pushInt(l * r)
	default:
		return fmt.Errorf("unknown operator at %s", op)
	}

	return nil
}
