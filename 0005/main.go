package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	lenS := len(s)
	maxLen := 0
	result := ""
	for start := range s {
		for finish := lenS; finish > start; finish-- {
			word := s[start:finish]
			palFlag := palindromCheck(word) == true
			if palFlag && len(word) > maxLen {
				maxLen = len(word)
				result = word
				if start < len(s)-1 && len(word) > len(s[start+1:]) {
					return word
				}
			}
			if palFlag {
				break
			}
		}
	}
	return result
}

func palindromCheck(s string) bool {
	lenS := len(s)
	if lenS == 1 {
		return true
	}
	middleOne := 0
	if lenS%2 == 1 {
		middleOne = 1
	}
	leftStraightPart := s[:lenS/2]
	tmp := s[lenS/2+middleOne:]
	rightReverserdPart := ""
	for i := len(tmp) - 1; i >= 0; i-- {
		rightReverserdPart = rightReverserdPart + string(tmp[i])
	}
	return leftStraightPart == rightReverserdPart

}

func main() {
	fmt.Println(longestPalindrome("aaxa"))
}
