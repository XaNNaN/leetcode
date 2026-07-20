package main

func shiftGrid(grid [][]int, k int) [][]int {
	rowsNum := len(grid)
	colsNum := len(grid[0])
	totalLen := rowsNum * colsNum
	result := make([][]int, rowsNum)
	for _, row := range result {
		row = make([]int, colsNum)
	}
	for i := range result {
		for j := range result[i] {
			idx := (i*rowsNum + j + k) % totalLen
			iNew := idx / colsNum
			jNew := idx % colsNum
			result[iNew][jNew] = grid[i][j]
		}
	}
	return result
}
