package main

import (
	"reflect"
	"testing"
)

func TestMinAbsDifference(t *testing.T) {
	var tests = []struct {
		input []int
		want  [][]int
	}{
		{[]int{4, 2, 1, 3}, [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{[]int{1, 3, 6, 10, 15}, [][]int{{1, 3}}},
		{[]int{3, 8, -10, 23, 19, -4, -14, 27}, [][]int{{-14, -10}, {19, 23}, {23, 27}}},
	}
	for _, test := range tests {
		if got := minimumAbsDifference(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("minimumAbsDifference(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

func BenchmarkMinimumAbsDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minimumAbsDifference([]int{ -467973, 725183, -256709, 879368, -246202, -339355, 151349})
	}
}
