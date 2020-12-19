package report

import (
	"reflect"
	"strings"
	"testing"
)

func TestCombinations(t *testing.T) {
	in := `1
2
3
4
`

	r, err := NewReport(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		size int
		want [][]int
	}{
		{
			0,
			[][]int{},
		},
		{
			1,
			[][]int{{1}, {2}, {3}, {4}},
		},
		{
			2,
			[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
	}

	for _, tc := range testCases {
		got := r.Combinations(tc.size)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("r.Combinations(%#v): got %#v, want %#v", tc.size, got, tc.want)
		}
	}
}

func TestProductOfCombinationWithSum(t *testing.T) {
	in := `1721
979
366
299
675
1456
`

	r, err := NewReport(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		size int
		want int
	}{
		{
			2,
			514579,
		},
		{
			3,
			241861950,
		},
	}

	for _, tc := range testCases {
		got := r.ProductOfCombinationWithSum(2020, tc.size)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("r.ProductOfCombinationWithSum(2020, %#v): got %#v, want %#v", tc.size, got, tc.want)
		}
	}
}
