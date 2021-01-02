package adapter

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadAdapters(t *testing.T) {
	in := `16
10
15
5
1
11
7
19
6
12
4
`

	got, err := ReadAdapters(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	want := Adapters{
		{joltage:1},
		{joltage:4},
		{joltage:5},
		{joltage:6},
		{joltage:7},
		{joltage:10},
		{joltage:11},
		{joltage:12},
		{joltage:15},
		{joltage:16},
		{joltage:19},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestCountAdapters(t *testing.T) {
	testCases := []struct {
		adapters Adapters
		want     map[int]int
	}{
		{
			Adapters{
				{joltage: 1},
				{joltage: 4},
				{joltage: 5},
				{joltage: 6},
				{joltage: 7},
				{joltage: 10},
				{joltage: 11},
				{joltage: 12},
				{joltage: 15},
				{joltage: 16},
				{joltage: 19},
			},
			map[int]int{1: 7, 3: 5},
		},
		{
			Adapters{
				{joltage: 1},
				{joltage: 2},
				{joltage: 3},
				{joltage: 4},
				{joltage: 7},
				{joltage: 8},
				{joltage: 9},
				{joltage: 10},
				{joltage: 11},
				{joltage: 14},
				{joltage: 17},
				{joltage: 18},
				{joltage: 19},
				{joltage: 20},
				{joltage: 23},
				{joltage: 24},
				{joltage: 25},
				{joltage: 28},
				{joltage: 31},
				{joltage: 32},
				{joltage: 33},
				{joltage: 34},
				{joltage: 35},
				{joltage: 38},
				{joltage: 39},
				{joltage: 42},
				{joltage: 45},
				{joltage: 46},
				{joltage: 47},
				{joltage: 48},
				{joltage: 49},
			},
			map[int]int{1: 22, 3: 10},
		},
	}

	for _, tc := range testCases {
		got := tc.adapters.CountAdapters()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %#v, want %#v", got, tc.want)
		}
	}
}

func TestCountChains(t *testing.T) {
	testCases := []struct {
		adapters Adapters
		want     int
	}{
		{
			Adapters{
				{joltage: 1},
				{joltage: 4},
				{joltage: 5},
				{joltage: 6},
				{joltage: 7},
				{joltage: 10},
				{joltage: 11},
				{joltage: 12},
				{joltage: 15},
				{joltage: 16},
				{joltage: 19},
			},
			8,
		},
		{
			Adapters{
				{joltage: 1},
				{joltage: 2},
				{joltage: 3},
				{joltage: 4},
				{joltage: 7},
				{joltage: 8},
				{joltage: 9},
				{joltage: 10},
				{joltage: 11},
				{joltage: 14},
				{joltage: 17},
				{joltage: 18},
				{joltage: 19},
				{joltage: 20},
				{joltage: 23},
				{joltage: 24},
				{joltage: 25},
				{joltage: 28},
				{joltage: 31},
				{joltage: 32},
				{joltage: 33},
				{joltage: 34},
				{joltage: 35},
				{joltage: 38},
				{joltage: 39},
				{joltage: 42},
				{joltage: 45},
				{joltage: 46},
				{joltage: 47},
				{joltage: 48},
				{joltage: 49},
			},
			19208,
		},
	}

	for _, tc := range testCases {
		a := Adapter{}
		got := a.CountChains(tc.adapters)
		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}
