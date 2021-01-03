package advent

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type schedule struct {
	index []int
	ids   map[int]int
}

func (s *schedule) read(rd io.Reader) error {
	var str string

	if _, err := fmt.Fscanln(rd, &str); err != nil && err != io.EOF {
		return err
	}

	buses := strings.Split(str, ",")

	s.index = make([]int, 0, len(buses))
	s.ids = make(map[int]int, len(buses))

	for i, f := range buses {
		if f == "x" {
			continue
		}

		id, err := strconv.Atoi(f)
		if err != nil {
			return err
		}

		s.index = append(s.index, i)
		s.ids[i] = id
	}

	return nil
}

func (s schedule) each(f func(i, id int)) {
	for _, i := range s.index {
		f(i, s.ids[i])
	}
}
