package main

import "fmt"

func main() {
	fmt.Println(countPartitions([]int{10, 10, 3, 7, 6}))
}

func countPartitions(nums []int) int {
	// Чётное +- Чётное = Чётное
	// Нечётное +- Нечётное = Чётное
	// Чётное +- Нечётное = Нечётное

	leftSum := 0
	rightSum := 0
	sum := 0
	counter := 0
	for _, n := range nums {
		sum += n
	}

	rightSum = sum
	for _, n := range nums[:len(nums)-1] {
		leftSum += n
		rightSum -= n
		if (leftSum-rightSum)%2 == 0 {
			counter++
		}
	}
	return counter
}
