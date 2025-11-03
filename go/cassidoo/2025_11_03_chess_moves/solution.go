package chessmoves

var kingMovePatterns = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

var knightMovePatterns = [][]int{
	{-2, -1},
	{-2, 1},
	{-1, -2},
	{-1, 2},
	{1, -2},
	{1, 2},
	{2, -1},
	{2, 1},
}

func KingMoves(currentPosition Position) []Position {
	positions := []Position{}
	for _, pattern := range kingMovePatterns {
		newPosition := generateNewPosition(currentPosition, pattern[0], pattern[1])
		if isInRange(newPosition) {
			positions = append(positions, newPosition)
		}
	}
	return positions
}

func KnightMoves(currentPosition Position) []Position {
	positions := []Position{}
	for _, pattern := range knightMovePatterns {
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
