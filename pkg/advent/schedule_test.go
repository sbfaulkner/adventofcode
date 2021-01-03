package advent

import (
	"strings"
	"testing"
)

func TestReadSchedule(t *testing.T) {
	in := `939
7,13,x,x,59,x,31,19
`

	s, err := ReadSchedule(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	t.Error("TODO", s)
}
