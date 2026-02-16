package zoom

import "slices"

func Zoom(grid [][]int, k int) [][]int {
	var zoomedGrid [][]int
	for _, row := range grid {
		var zoomedRow []int
		for _, item := range row {
			zoomedRow = append(zoomedRow, slices.Repeat([]int{item}, k)...)
		}

		for range k {
			zoomedGrid = append(zoomedGrid, zoomedRow)
		}
	}
	return zoomedGrid
}
