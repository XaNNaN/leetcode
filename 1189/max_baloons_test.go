package maxballoons

import (
	"fmt"
	"testing"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	var tests = []struct {
		arg  string
		want int
	}{
		{"nlaebolko", 1},
		{"loonbalxballpoon", 2},
		{"leetcode", 0},
		{"nllbblooon", 0},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v, %d)", test.arg, test.want)
		got := MaxNumberOfBalloons(test.arg)
		if got != test.want {
			t.Errorf("%s = %v, требуется %v", descr, got, test.want)
		}
	}
}
