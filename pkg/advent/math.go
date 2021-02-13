package advent

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func (t token) String() string {
	switch t.t {
	case eofToken:
		return "<EOF>"
	case intToken:
		return fmt.Sprintf("<%d>", t.i)
	default:
		return fmt.Sprintf("<%s>", t.s)
	}
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

// Evaluate solves a math problem
func (p *MathProblem) Evaluate(fp OperatorPrecedence) (int, error) {
	e := mathEvaluator{
		rd:          strings.NewReader(string(*p)),
		fprecedence: fp,
	}

	fstate := evaluateLeftValue

	for e.getToken() {
		fs, err := fstate(&e)
		if err != nil {
			return 0, err
		}
		fstate = fs
	}

	i := e.popToken().i

	return i, nil
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

type mathEvaluator struct {
	rd          *strings.Reader
	err         error
	depth       int
	tok         token
	unget       bool
	st          []token
	fstate      mathEvaluatorState
	fprecedence OperatorPrecedence
}

func (e *mathEvaluator) getToken() bool {
	if e.unget {
		e.unget = false
		return true
	}

	for {
		ch, _, err := e.rd.ReadRune()
		if err == io.EOF {
			e.tok.t = eofToken
			return len(e.st) > 1
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

func (e *mathEvaluator) ungetToken() {
	if e.unget {
		log.Fatal("unable to unget multiple tokens")
	}

	if e.tok.t != eofToken {
		e.unget = true
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

func (e *mathEvaluator) pushToken(t token) {
	e.st = append(e.st, t)
}

func (e *mathEvaluator) pushInt(i int) {
	e.pushToken(token{t: intToken, i: i})
}

func (e *mathEvaluator) popToken() token {
	if len(e.st) == 0 {
		return token{t: eofToken}
	}

	t := e.st[len(e.st)-1]
	e.st = e.st[0 : len(e.st)-1]
	return t
}

type mathEvaluatorState func(*mathEvaluator) (mathEvaluatorState, error)

func evaluateLeftValue(e *mathEvaluator) (mathEvaluatorState, error) {
	switch e.tok.t {
	case intToken:
		e.pushToken(e.tok)
		return evaluateOperator, nil

	case lparenToken:
		e.pushToken(e.tok)
		return evaluateLeftValue, nil
	}

	return nil, fmt.Errorf("syntax error at %s (expected integer)", e.tok.s)
}

func evaluateOperator(e *mathEvaluator) (mathEvaluatorState, error) {
	if e.tok.t == rparenToken {
		i := e.popToken()
		e.popToken()

		e.pushToken(i)

		return evaluateOperator, nil
	}

	if e.tok.t != opToken {
		return nil, fmt.Errorf("syntax error at %s (expected operator)", e.tok.s)
	}

	e.pushToken(e.tok)

	return evaluateRightValue, nil
}

func evaluateRightValue(e *mathEvaluator) (mathEvaluatorState, error) {
	switch e.tok.t {
	case intToken:
		e.pushToken(e.tok)
		return evaluateExpression, nil

	case lparenToken:
		e.pushToken(e.tok)
		return evaluateLeftValue, nil
	}

	return nil, fmt.Errorf("syntax error at %s (expected rval)", e.tok.s)
}

func evaluateExpression(e *mathEvaluator) (mathEvaluatorState, error) {
	r := e.popToken()
	op := e.popToken()

	if e.tok.t == rparenToken && op.t == lparenToken {
		e.pushToken(r)
		return evaluateExpression, nil
	}

	if op.t != opToken {
		if op.t != eofToken {
			e.pushToken(op)
		}
		e.pushToken(r)
		e.ungetToken()

		return evaluateOperator, nil
	}

	if e.tok.t == opToken && !e.fprecedence(op.s, e.tok.s) {
		e.pushToken(op)
		e.pushToken(r)
		e.pushToken(e.tok)

		return evaluateRightValue, nil
	}

	if e.tok.t != eofToken {
		e.ungetToken()
	}

	l := e.popToken()

	switch op.s {
	case "+":
		e.pushInt(l.i + r.i)
	case "*":
		e.pushInt(l.i * r.i)
	default:
		return nil, fmt.Errorf("unknown operator at %s", op)
	}

	return evaluateExpression, nil
}
