package chessmoves

var KingMoves = movesGenerator([][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
})

var KnightMoves = movesGenerator([][]int{
	{-2, -1},
	{-2, 1},
	{-1, -2},
	{-1, 2},
	{1, -2},
	{1, 2},
	{2, -1},
	{2, 1},
})

func movesGenerator(patterns [][]int) func(Position) []Position {
	return func(currentPosition Position) []Position {
		return generateMoves(currentPosition, patterns)
	}
}

func generateMoves(currentPosition Position, patterns [][]int) []Position {
	positions := []Position{}
	for _, pattern := range patterns {
		newPosition := generateNewPosition(currentPosition, pattern[0], pattern[1])
		if isInRange(newPosition) {
			positions = append(positions, newPosition)
		}
	}
	return positions
}

func generateNewPosition(currentPosition Position, rowChange, colChange int) Position {
	return Position{row: currentPosition.row + rowChange, col: currentPosition.col + colChange}
}

func isInRange(position Position) bool {
	return position.row >= 0 && position.row <= 7 && position.col >= 0 && position.col <= 7
}
