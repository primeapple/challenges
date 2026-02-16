package zoom

import (
	"cassidoo/utils"
	"testing"
)

func TestZoomAI(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := ZoomAI([][]int{{1, 2}, {3, 4}}, 2)
		want := [][]int{
			{1, 1, 2, 2},
			{1, 1, 2, 2},
			{3, 3, 4, 4},
			{3, 3, 4, 4},
		}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := ZoomAI([][]int{{7, 8, 9}}, 3)
		want := [][]int{
			{7, 7, 7, 8, 8, 8, 9, 9, 9},
			{7, 7, 7, 8, 8, 8, 9, 9, 9},
			{7, 7, 7, 8, 8, 8, 9, 9, 9},
		}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := ZoomAI([][]int{{1}, {2}}, 3)
		want := [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
			{2, 2, 2},
			{2, 2, 2},
			{2, 2, 2},
		}
		utils.AssertDeepEquals(t, got, want)
	})

}
