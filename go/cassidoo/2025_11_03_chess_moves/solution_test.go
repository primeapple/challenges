package chessmoves

import (
	"cassidoo/utils"
	"testing"
)

func TestKingMoves(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		position := Position{row: 4, col: 4}

		got := KingMoves(position)
		want := []Position{
			{row: 3, col: 3},
			{row: 3, col: 4},
			{row: 3, col: 5},
			{row: 4, col: 3},
			{row: 4, col: 5},
			{row: 5, col: 3},
			{row: 5, col: 4},
			{row: 5, col: 5},
		}

		utils.AssertDeepEquals(t, got, want)
	})
}

func TestKnightMoves(t *testing.T) {
	// EXAMPLES
	t.Run("should work with example 1", func(t *testing.T) {
		position := Position{row: 4, col: 4}

		got := KnightMoves(position)
		want := []Position{
			{row: 2, col: 3},
			{row: 2, col: 5},
			{row: 3, col: 2},
			{row: 3, col: 6},
			{row: 5, col: 2},
			{row: 5, col: 6},
			{row: 6, col: 3},
			{row: 6, col: 5},
		}

		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		position := Position{row: 0, col: 0}

		got := KnightMoves(position)
		want := []Position{
			{row: 1, col: 2},
			{row: 2, col: 1},
		}

		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		position := Position{row: 1, col: 2}

		got := KnightMoves(position)
		want := []Position{
			{row: 0, col: 0},
			{row: 0, col: 4},
			{row: 2, col: 0},
			{row: 2, col: 4},
			{row: 3, col: 1},
			{row: 3, col: 3},
		}

		utils.AssertDeepEquals(t, got, want)
	})
}
