package report

import (
	"bufio"
	"io"
	"math/big"
	"strconv"
)

// Report is an expense report
type Report []int

// NewReport creates a new Report loading it from a reader
func NewReport(rd io.Reader) (*Report, error) {
	r := Report{}

	err := r.read(rd)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *Report) read(rd io.Reader) error {
	s := bufio.NewScanner(rd)

	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			return err
		}

		*r = append(*r, i)
	}

	return s.Err()
}

// Combinations gets the combinations of the specified size
func (r Report) Combinations(choose int) [][]int {
	return combinations(r, choose)
}

func combinations(set []int, choose int) [][]int {
	if choose == 0 {
		return [][]int{}
	}

	c := make([][]int, 0, numberOfCombinations(len(set), choose))

	for i, n := range set[:len(set)-choose+1] {
		a := []int{n}

		if choose == 1 {
			c = append(c, a)
			continue
		}

		b := combinations(set[i+1:], choose-1)

		for _, r := range b {
			c = append(c, append(a, r...))
		}
	}

	return c
}

func numberOfCombinations(n, r int) int {
	if r == n {
		return 1
	}

	var p, d big.Int

	d.Div(factorial(n), p.Mul(factorial(r), factorial(n-r)))

	return int(d.Int64())
}

func factorial(n int) *big.Int {
	if n == 1 {
		return big.NewInt(1)
	}

	var f big.Int

	return f.Mul(big.NewInt(int64(n)), factorial(n-1))
}

// ProductOfCombinationWithSum calculates the product of a combination of the specified size with the specified sum
func (r Report) ProductOfCombinationWithSum(sum, choose int) int {
	for _, c := range r.Combinations(choose) {
		s := 0
		p := 1

		for _, n := range c {
			s += n
			p *= n
		}

		if s == sum {
			return p
		}
	}

	return -1
}
