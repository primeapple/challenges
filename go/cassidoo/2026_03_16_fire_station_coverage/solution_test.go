package fireStationCoverage

import (
	"cassidoo/utils"
	"testing"
)

func TestFireStationCoverage(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := FireStationCoverage([][]CellStatus{
			{2, 0, 1},
			{0, 2, 0},
			{1, 0, 2},
		})
		want := [][]int{
			{2, 1, 0},
			{1, 2, 1},
			{0, 1, 2},
		}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := FireStationCoverage([][]CellStatus{
			{1, 0, 0, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 0, 0, 1},
		})
		want := [][]int{
			{0, 1, 1, 0},
			{1, 2, 2, 1},
			{1, 2, 2, 1},
			{0, 1, 1, 0},
		}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work only fire hydrants", func(t *testing.T) {
		got := FireStationCoverage([][]CellStatus{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		})
		want := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		utils.AssertDeepEquals(t, got, want)
	})
}
