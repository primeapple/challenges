package chessmoves

func KnightMovesAI(position Position) []Position {
	moves := []Position{}

	knightOffsets := [][]int{
		{-2, -1}, {-2, 1},
		{-1, -2}, {-1, 2},
		{1, -2}, {1, 2},
		{2, -1}, {2, 1},
	}

	for _, offset := range knightOffsets {
		newRow := position.row + offset[0]
		newCol := position.col + offset[1]

		if newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8 {
			moves = append(moves, Position{row: newRow, col: newCol})
		}
	}

	return moves
}

func KingMovesAI(position Position) []Position {
	moves := []Position{}

	kingOffsets := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, offset := range kingOffsets {
		newRow := position.row + offset[0]
		newCol := position.col + offset[1]

		if newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8 {
			moves = append(moves, Position{row: newRow, col: newCol})
		}
	}

	return moves
}
