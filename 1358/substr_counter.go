package task1358

type letterInfo struct {
	count     int
	positions []int
}

func numberOfSubstrings(s string) int {
	letters := make(map[string]letterInfo)
	letters["a"] = letterInfo{0, make([]int, 0)}
	letters["b"] = letterInfo{0, make([]int, 0)}
	letters["c"] = letterInfo{0, make([]int, 0)}
	for _, char := range s {
		switch string(char) {
		case "a":
			entry := letters["a"]
			entry.count++
			letters["a"] = entry
		case "b":
			entry := letters["b"]
			entry.count++
			letters["b"] = entry
		case "c":
			entry := letters["c"]
			entry.count++
			letters["c"] = entry
		}
	}
	// TODO: add tails. Valid substr which starts and ends with the same letter. Depends on positions???
	return letters["a"].count * letters["b"].count * letters["c"].count
}
