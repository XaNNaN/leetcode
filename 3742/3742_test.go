package main

import (
	"testing"
)

func TestSimpleCorrect(t *testing.T) {
	wanted := 2
	grid := [][]int{{0, 1}, {1, 2}}
	k := 1
	score := maxPathScore(grid, k)
	if score != wanted {
		t.Errorf(`maxPathScore(%v) = %v, want %v`, grid, score, wanted)
	}
}
