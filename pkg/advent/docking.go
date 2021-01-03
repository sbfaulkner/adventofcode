package advent

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// DockingProgram is a ferry docking program
type DockingProgram map[int]int64

// Initialize a ferry DockingProgram
func (dp *DockingProgram) Initialize(rd io.Reader) error {
	s := bufio.NewScanner(rd)

	maskReplacer := strings.NewReplacer("X", "1")
	bitReplacer := strings.NewReplacer("X", "0")

	var mask, bits int64
	var err error

	for s.Scan() {
		var dst, src string

		fmt.Sscanf(s.Text(), "%s = %s", &dst, &src)

		if dst == "mask" {
			mask, err = strconv.ParseInt(maskReplacer.Replace(src), 2, 64)
			if err != nil {
				return err
			}

			bits, err = strconv.ParseInt(bitReplacer.Replace(src), 2, 64)
			if err != nil {
				return err
			}

			continue
		}

		var addr int

		fmt.Sscanf(dst, "mem[%d]", &addr)

		var val int64

		val, err = strconv.ParseInt(src, 10, 64)
		if err != nil {
			return err
		}

		val &= mask
		val |= bits

		(*dp)[addr] = val
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}

// CheckSum returns the sum of all initialized memory locations
func (dp DockingProgram) CheckSum() uint64 {
	var sum uint64

	for _, val := range dp {
		sum += uint64(val)
	}

	return sum
}
