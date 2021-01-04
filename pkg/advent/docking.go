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
	set(dp *DockingProgram, addr int64, val int64) error
}

// DecoderV1 implements version 1 decoder chip
type DecoderV1 struct {
	mask int64
	bits int64
}

func (d *DecoderV1) setMask(src string) error {
	var err error

	maskReplacer := strings.NewReplacer("X", "1")
	bitReplacer := strings.NewReplacer("X", "0")

	d.mask, err = strconv.ParseInt(maskReplacer.Replace(src), 2, 64)
	if err != nil {
		return err
	}

	d.bits, err = strconv.ParseInt(bitReplacer.Replace(src), 2, 64)

	return err
}

func (d DecoderV1) set(dp *DockingProgram, addr int64, val int64) error {
	val &= d.mask
	val |= d.bits

	(*dp)[addr] = val

	return nil
}

// DecoderV2 implements version 2 decoder chip
type DecoderV2 struct {
	mask  int64
	bits  string
	count int
}

func pow2(i int) int {
	p := 1

	for i > 0 {
		p *= 2
		i--
	}

	return p
}

func (d *DecoderV2) setMask(src string) error {
	var err error

	maskReplacer := strings.NewReplacer("0", "1", "X", "0")

	d.mask, err = strconv.ParseInt(maskReplacer.Replace(src), 2, 64)
	if err != nil {
		return err
	}

	d.bits = src
	d.count = strings.Count(src, "X")

	return nil
}

func (d DecoderV2) mapFloatingBits(i int) func(rune) rune {
	var b int

	bits := []rune(fmt.Sprintf("%0*b", d.count, i))

	return func(r rune) rune {
		if r == 'X' {
			r = bits[b]
			b++
		}
		return r
	}
}

func (d DecoderV2) set(dp *DockingProgram, addr int64, val int64) error {
	addr &= d.mask

	for i := 0; i < pow2(d.count); i++ {
		bits, err := strconv.ParseInt(strings.Map(d.mapFloatingBits(i), d.bits), 2, 64)
		if err != nil {
			return err
		}

		a := addr | bits

		(*dp)[a] = val
	}

	return nil
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
			if err := decoder.setMask(src); err != nil {
				return err
			}
			continue
		}

		var addr int64

		fmt.Sscanf(dst, "mem[%d]", &addr)

		var val int64

		val, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			return err
		}

		err = decoder.set(dp, addr, val)
		if err != nil {
			return err
		}
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
