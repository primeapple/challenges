package latinSquare

import (
	"slices"
)

func LatinSquare(n int) [][]int {
	result := make([][]int, 0, n)

	initialRow := make([]int, 0, n)
	for i := range n {
		initialRow = append(initialRow, i+1)
	}
	result = append(result, initialRow)

	for i := 1; i < n; i++ {
		result = append(result, rotateBy(initialRow, i))
	}

	return result
}

func LatinSquareConcurrent(n int) [][]int {
	result := make([][]int, 0, n)

	initialRow := make([]int, 0, n)
	for i := range n {
		initialRow = append(initialRow, i+1)
	}
	result = append(result, initialRow)

	c := make(chan []int, n)
	done := make(chan struct{}, n)

	for i := 1; i < n; i++ {
		go rotateByWithChannel(initialRow, i, c, done)
	}
	for i := 1; i < n; i++ {
		<-done
	}
	close(c)

	for i := 1; i < n; i++ {
		result = append(result, <-c)
	}
	slices.SortFunc(result, func(a, b []int) int {
		return a[0] - b[0]
	})

	return result
}

func rotateBy(row []int, offset int) []int {
	beginning := row[offset:]
	end := row[0:offset]
	return append(beginning, end...)
}

func rotateByWithChannel(row []int, offset int, c chan []int, done chan struct{}) {
	beginning := row[offset:]
	end := row[0:offset]
	c <- append(beginning, end...)
	done <- struct{}{}
}
