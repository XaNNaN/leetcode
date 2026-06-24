package pascal_triangle
// Just a comment

func generate(numRows int) [][]int {
	rowLen := 0
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		rowLen += 1
		result[i] = make([]int, rowLen)
		for y:= 0; y < rowLen; y++ {
			parentLeft :=  0
			if rowLen - 1 > 0 && y - 1 >= 0 {
				parentLeft = result[i - 1][y - 1]
			}
			parentRight := 0
			if rowLen - 1 > 0 && y < rowLen - 1 {
				parentRight = result[i - 1][y]
			}
			if rowLen - 1 == 0 {
				parentLeft = 1
			} 
			result[i][y] = parentLeft + parentRight
		}
	}
	return result
}
