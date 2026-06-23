package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	result := make([][]int, 0, len(arr))
	minDiff := math.Inf(1)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			curDiff := math.Abs(float64(arr[i] - arr[j]))
			if curDiff < minDiff {
				minDiff = curDiff
				result = [][]int{{arr[i], arr[j]}}
			} else if curDiff == minDiff {
				result = append(result, []int{arr[i], arr[j]})
			}
		}
	}
	//fmt.Printf("Minimum Difference: %v\n", minDiff)
	//fmt.Printf("Diff map: %v\n", diffMap)

	return result
}
