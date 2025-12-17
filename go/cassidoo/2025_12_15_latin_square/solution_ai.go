package latinSquare

import "sync"

func LatinSquareAI(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = (i+j)%n + 1
		}
	}
	return result
}

func LatinSquareConcurrentAI(n int) [][]int {
	result := make([][]int, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			result[row] = make([]int, n)
			for j := 0; j < n; j++ {
				result[row][j] = (row+j)%n + 1
			}
		}(i)
	}
	wg.Wait()
	return result
}
