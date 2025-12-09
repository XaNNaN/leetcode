package main

import "math"

func main() {
	countPartitions([]int{9, 4, 1, 3, 7}, 4)
}

func countPartitions(nums []int, k int) int {
	//partition
	// contiguous segments
	// abs(max - min) <= k
	// return modulo 10^9 + 7
	result := 0
	rightMax := 0
	for _, num := range nums {
		rightMax += num
	}
	leftSum := 0
	for _, num := range nums[:len(nums)-1] {
		leftSum += num
		rightMax -= num
		if int(math.Abs(float64(leftSum-rightMax))) <= k {
			result++
		}
	}

	return result % (1000000000 + 7)
}
