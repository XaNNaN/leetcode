package _211

import "fmt"

func main() {
	fmt.Println(countCollisions("RLRSLL")) // 2 2 S 1 1
	fmt.Println(countCollisions("LLRR"))
	fmt.Println(countCollisions("LLRLRLLSLRLLSLSSSS"))
	fmt.Println(countCollisions("SSRSSRLLRSLLRSRSSRLRRRRLLRRLSSRR"))
	// LLSSS
}

func countCollisions(directions string) int {
	// endless road
	// car>car><car
	// n cars
	// each car present at a unique point
	// directions[i]:
	// "L" - left
	// "R" - right
	// "S" - stop
	// Get the total collision counter
	// RL -> colCounter += 2
	// RS or SL -> colCounter++

	colCounter := 0
	rightCounter := 0
	carMap := make(map[int]string)
	for i, dir := range directions {
		carMap[i] = string(dir)
	}

	for i, _ := range directions[:len(directions)-1] {
		switch ev := carMap[i] + carMap[i+1]; ev {
		case "RL":
			colCounter += 2 + rightCounter
			rightCounter = 0
			carMap[i] = "S"
			carMap[i+1] = "S"
		case "RS":
			colCounter += 1 + rightCounter
			rightCounter = 0
			carMap[i] = "S"
		case "SL":
			colCounter += 1
			carMap[i+1] = "S"
		case "RR":
			rightCounter += 1

		}
	}
	if carMap[len(directions)-1] == "L" && carMap[len(directions)-2] == "S" {
		colCounter += 1
		carMap[len(directions)-1] = "S"
	}
	fmt.Println(carMap, colCounter)
	return colCounter
}
