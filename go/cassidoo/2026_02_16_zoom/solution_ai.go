package zoom

// ZoomAI zooms in on a 2D grid by factor k, turning each cell into a k x k block
func ZoomAI(grid [][]int, k int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return [][]int{}
	}

	rows := len(grid)
	cols := len(grid[0])

	// Result dimensions: each original cell becomes k x k cells
	newRows := rows * k
	newCols := cols * k

	result := make([][]int, newRows)
	for i := range result {
		result[i] = make([]int, newCols)
	}

	// For each cell in the original grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			value := grid[i][j]

			// Fill the k x k block in the result grid
			for di := 0; di < k; di++ {
				for dj := 0; dj < k; dj++ {
					result[i*k+di][j*k+dj] = value
				}
			}
		}
	}

	return result
}
