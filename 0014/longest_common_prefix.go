// Package prefix returns LongestCommonPrefix
package prefix

import "math"

func LongestCommonPrefix(strs []string) string {
	// max len of the longest prefix will be equal to len of the shortest str.
	// 1. Find shortest string
	// 2. Use it as a base solution
	// 3. Split str solution := solution[: len(solution)/2]
	// 4.
	var step, prefixLen int
	minLen := math.MaxInt64
	for _, str := range strs {
		strLen := len(str)
		if strLen < minLen {
			minLen = strLen
		}
	}
	minLen0 := minLen
	step = minLen / 2
	minLen = minLen - step
	lowBorder := 0
	topBorder := minLen0
	for step >= 0 {
		equal := true
		prevStr := strs[0]
		for _, str := range strs {
			if str[:minLen] != prevStr[:minLen] {
				equal = false
				break
			}
		}
		if equal {
			// topBorder = minLen0
			lowBorder = minLen
			step = (topBorder - lowBorder) / 2
			if topBorder-lowBorder == 1 && topBorder == minLen0 && minLen0 != 1 {
				step = 1
			}
			prefixLen = minLen
			minLen = minLen + step
		}
		if !equal {
			topBorder = minLen
			step = (topBorder - lowBorder) / 2
			if topBorder-lowBorder == 1 && lowBorder == 0 && minLen0 != 1 {
				step = 1
			}
			minLen = minLen - step
		}
		if step == 0 {
			break
		}
	}
	return strs[0][:prefixLen]
}
