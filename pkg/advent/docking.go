package advent

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Decoder interface implemented by decoder chip for ferry docking program
type Decoder interface {
	setMask(s string) error
	set(dp *DockingProgram, addr int64, val int64)
}

// DecoderV1 implements version 1 decoder chip
type DecoderV1 struct {
	mask int64
	bits int64
}

func (d *DecoderV1) setMask(src string) (err error) {
	maskReplacer := strings.NewReplacer("X", "1")
	bitReplacer := strings.NewReplacer("X", "0")

	d.mask, err = strconv.ParseInt(maskReplacer.Replace(src), 2, 64)
	if err == nil {
		d.bits, err = strconv.ParseInt(bitReplacer.Replace(src), 2, 64)
	}

	return
}

func (d DecoderV1) set(dp *DockingProgram, addr int64, val int64) {
	val &= d.mask
	val |= d.bits

	(*dp)[addr] = val
}

// DockingProgram is a ferry docking program
type DockingProgram map[int64]int64

// Initialize a ferry DockingProgram
func (dp *DockingProgram) Initialize(rd io.Reader, decoder Decoder) error {
	s := bufio.NewScanner(rd)

	for s.Scan() {
		var dst, src string

		fmt.Sscanf(s.Text(), "%s = %s", &dst, &src)

		if dst == "mask" {
			decoder.setMask(src)
			continue
		}

		var addr int64

		fmt.Sscanf(dst, "mem[%d]", &addr)

		var val int64

		val, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			return err
		}

		decoder.set(dp, addr, val)
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
