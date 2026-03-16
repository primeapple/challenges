package fireStationCoverage

import (
	"math"
)

func FireStationCoverage(city [][]CellStatus) [][]int {
	distances := make([][]int, len(city))

	for y, cityRow := range city {
		distancesRow := make([]int, len(cityRow))
		for x, cityCell := range cityRow {
			if cityCell == FireStation {
				distancesRow[x] = 0
			} else {
				distancesRow[x] = -1
			}
		}
		distances[y] = distancesRow
	}

	hasUpdated := true
	for ;hasUpdated; {
		hasUpdated = false
		for y, distancesRow := range distances {
			for x, distance := range distancesRow {
				minimalSurroundingDistance := getMinimalSurroundingDistance(y, x, distances)
				if minimalSurroundingDistance == -1 {
					continue
				}
				if distance != -1 && distance <= minimalSurroundingDistance + 1 {
					continue
				}
				distancesRow[x] = minimalSurroundingDistance + 1
				hasUpdated = true
			}
		}
	}

	return distances
}

func getMinimalSurroundingDistance(y, x int, distances [][]int) int {
	top := safeGet(x, y+1, distances)
	bottom := safeGet(x, y-1, distances)
	left := safeGet(x-1, y, distances)
	right := safeGet(x+1, y, distances)

	return minimalDistance(top, bottom, left, right)
}

func safeGet(x, y int, distances [][]int) int {
	if y < 0 || y >= len(distances) || x < 0 || x >= len(distances[y]) {
		return -1
	}
	return distances[y][x]
}

func minimalDistance(d1, d2, d3, d4 int) int {
	if d1 == -1 && d2 == -1 && d3 == -1 && d4 == -1 {
		return -1
	}

	minDistance := math.MaxInt
	if d1 != -1 && d1 < minDistance {
		minDistance = d1
	}
	if d2 != -1 && d2 < minDistance {
		minDistance = d2
	}
	if d3 != -1 && d3 < minDistance {
		minDistance = d3
	}
	if d4 != -1 && d4 < minDistance {
		minDistance = d4
	}

	return minDistance
}
