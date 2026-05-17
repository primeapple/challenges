package countBouncesToTarget

import (
	"fmt"
)

func CountBouncesToTarget(grid, start, target, direction []int) int {
	current := []int{start[0], start[1]}
	bounces := 0
	seen := map[string]struct{}{}

	for {
		if current[0] == target[0] && current[1] == target[1] {
			break
		}
		bounces = bounces + nextStep(grid, current, direction)
		key := fmt.Sprintf("%d-%d-%d-%d", current[0], current[1], direction[0], direction[1])
		_, alreadySeen := seen[key]
		if alreadySeen {
			return -1
		}
		seen[key] = struct{}{}
	}
	return bounces
}

func nextStep(grid, current, direction []int) int {
	newX, bouncesX := CalculateNextPosition(current[0], grid[0]-1, direction[0])
	newY, bouncesY := CalculateNextPosition(current[1], grid[1]-1, direction[1])
	if bouncesX % 2 == 1 {
		direction[0] = -direction[0]
	}
	if bouncesY % 2 == 1 {
		direction[1] = -direction[1]
	}
	println(newX, newY, direction[0], direction[1])
	current[0], current[1] = newX, newY
	return bouncesX + bouncesY
}

func CalculateNextPosition(current, maxSize, direction int) (int, int) {
	nextPosition := current + direction
	if nextPosition > maxSize {
		positionAfterBounces, bounces := CalculateNextPosition(maxSize, maxSize, -(direction - (maxSize - current)))
		return positionAfterBounces, bounces + 1
	}
	if nextPosition < 0 {
		positionAfterBounces, bounces := CalculateNextPosition(0, maxSize, -(direction - (0 - current)))
		return positionAfterBounces, bounces + 1
	}
	return nextPosition, 0
}
