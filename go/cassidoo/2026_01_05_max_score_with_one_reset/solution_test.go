package maxScoreWithOneReset

import (
	"cassidoo/utils"
	"testing"
)

func TestMaxScoreWithOneReset(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := MaxScoreWithOneReset([]int{2, -1, 2, -5, 2, 2})
		want := 4
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := MaxScoreWithOneReset([]int{4, -10, 3, 2, -1, 6})
		want := 10
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := MaxScoreWithOneReset([]int{-50, -2, -3})
		want := 0
		utils.AssertDeepEquals(t, got, want)
	})
}
