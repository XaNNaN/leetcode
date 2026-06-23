package maxballoons

import "fmt"

func MaxNumberOfBalloons(text string) int {
	var b, a, l, o, n int
	for _, letter := range text {
		switch string(letter) {
		case "b":
			b += 1
		case "a":
			a += 1
		case "l":
			l += 1
		case "o":
			o += 1
		case "n":
			n += 1
		}
	}
	l = l / 2
	o = o / 2
	min := b
	for _, k := range []int{a, l, o, n} {
		if k < min {
			min = k
		}
	}
	fmt.Printf("b: %v, a: %v, l:%v, o: %v, n: %v\n", b, a, l, o, n)
	return min
}
