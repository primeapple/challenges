package minRepairs

import (
	"fmt"
	"maps"
	"slices"
)

// Idea:
// Divide and conquer: Identify unique broken regions
// For each, find the minimum repairs needed via brute force, then summarize
func MinRepairs(grid [][]int, k int) int {
	brokenComponents := FindBrokenComponents(grid)

	repairsNeeded := 0
	for _, brokenComponent := range brokenComponents {
		if len(brokenComponent) <= k {
			continue
		}
		repairsNeeded += minRepairsNeeded(brokenComponent, k)
	}
	return repairsNeeded
}

func FindBrokenComponents(grid [][]int) []Component {
	components := map[int]Component{}

	// let's color the components
	currrentColor := 2
	for y, row := range grid {
		for x, cell := range row {
			if cell != 0 {
				continue
			}
			colorComponent(y, x, grid, currrentColor)
			currrentColor++
		}
	}

	// let's find the colored components
	for y, row := range grid {
		for x, cell := range row {
			if cell < 2 {
				continue
			}
			components[cell-2] = append(components[cell-2], Coordinate{y, x})
		}
	}

	return slices.Collect(maps.Values(components))
}

func colorComponent(y, x int, grid [][]int, color int) {
	if y < 0 || x < 0 || y >= len(grid) || x >= len(grid[y]) || grid[y][x] != 0 {
		return
	}
	grid[y][x] = color
	colorComponent(y-1, x, grid, color)
	colorComponent(y, x-1, grid, color)
	colorComponent(y+1, x, grid, color)
	colorComponent(y, x+1, grid, color)
}

func minRepairsNeeded(component Component, k int) int {
	if len(component) <= k {
		return 0
	}

	for numRepairs := 1; numRepairs < len(component); numRepairs++ {
		combos := combinations(component, numRepairs)
		for _, combo := range combos {
			removeSet := make(map[string]bool)
			for _, c := range combo {
				removeSet[coordKey(c)] = true
			}

			// Build remaining cells
			remaining := Component{}
			for _, c := range component {
				if !removeSet[coordKey(c)] {
					remaining = append(remaining, c)
				}
			}

			// Check if all sub-components are <= k
			if allComponentsSmallEnough(remaining, k) {
				return numRepairs
			}
		}
	}
	return len(component) // worst case: repair everything
}

// Generate all combinations of size r from the component
func combinations(component Component, r int) []Component {
	var result []Component
	var backtrack func(start int, current Component)
	backtrack = func(start int, current Component) {
		if len(current) == r {
			result = append(result, slices.Clone(current))
			return
		}
		for i := start; i < len(component); i++ {
			current = append(current, component[i])
			backtrack(i+1, current)
			current = current[:len(current)-1]
		}
	}
	backtrack(0, Component{})
	return result
}

// Check if all connected components in the remaining cells have size <= k
func allComponentsSmallEnough(cells Component, k int) bool {
	visited := make(map[string]bool)
	for _, cell := range cells {
		if visited[coordKey(cell)] {
			continue
		}

		// BFS to find component size
		size := 0
		queue := []Coordinate{cell}
		visited[coordKey(cell)] = true

		cellSet := make(map[string]bool)
		for _, c := range cells {
			cellSet[coordKey(c)] = true
		}

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			size++

			// Check all 4 neighbors
			for _, neighbor := range []Coordinate{{curr[0] - 1, curr[1]}, {curr[0] + 1, curr[1]}, {curr[0], curr[1] - 1}, {curr[0], curr[1] + 1}} {
				key := coordKey(neighbor)
				if cellSet[key] && !visited[key] {
					visited[key] = true
					queue = append(queue, neighbor)
				}
			}
		}

		if size > k {
			return false
		}
	}
	return true
}

func coordKey(c Coordinate) string {
	return fmt.Sprintf("%d-%d", c[0], c[1])
}
