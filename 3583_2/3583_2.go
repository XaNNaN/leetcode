package main

import "fmt"

func main() {
	fmt.Println(specialTriplets([]int{0, 1, 0, 0}))
	fmt.Println(specialTriplets([]int{8, 4, 2, 8, 4}))
}

// Решение через мапы с частотой каждого элемента до и после для каждого элемента
// Слайс мап инт инт
func specialTriplets(nums []int) int {
	result := 0

	freqPrev := make([]map[int]int, len(nums))
	freqNext := make([]map[int]int, len(nums))

	// init all maps
	for i, _ := range nums {
		freqPrev[i] = make(map[int]int)
		freqNext[i] = make(map[int]int)
	}

	// fill the first freqNext map
	for _, num := range nums[1:] {
		freqNext[0][num]++
	}

	for i, num := range nums[1:] {
		// копируем мапы с предыдущего индекса
		for key, val := range freqNext[i] {
			freqNext[i+1][key] = val
		}
		for key, val := range freqPrev[i] {
			freqPrev[i+1][key] = val
		}
		// Обновляем значения в новых мапах, убирая текущий элемент из мапы для последующих
		// и добавляя предыдущий элемент в мапу для предыдущих
		freqNext[i+1][num]--
		freqPrev[i+1][nums[i]]++
	}
	//fmt.Println("freqPrev", freqPrev)
	//fmt.Println("freqNext", freqNext)

	// Main cycle: running throw nums. Looking in freqPrev to find amount of num*2
	// and in freqNext to find amount of num*2 in freqNext

	for i, num := range nums {
		result += freqPrev[i][num*2] * freqNext[i][num*2]
	}

	return result % (1000000000 + 7)
}
