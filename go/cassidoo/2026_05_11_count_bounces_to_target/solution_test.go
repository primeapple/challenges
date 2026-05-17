package countBouncesToTarget

import (
	"cassidoo/utils"
	"testing"
)

func TestCountBouncesToTarget(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := CountBouncesToTarget([]int{8, 8}, []int{0, 0}, []int{3, 4}, []int{1, 4})
		want := -1
		utils.AssertEquals(t, got, want)
	})
	t.Run("should work with example 2", func(t *testing.T) {
		got := CountBouncesToTarget([]int{3, 3}, []int{0, 1}, []int{2, 1}, []int{1, 1})
		want := 1
		utils.AssertEquals(t, got, want)
	})
	t.Run("should work with example 3", func(t *testing.T) {
		got := CountBouncesToTarget([]int{4, 5}, []int{0, 0}, []int{3, 3}, []int{1, 1})
		want := 0
		utils.AssertEquals(t, got, want)
	})
}

func TestCalculateNextPosition(t *testing.T) {
	t.Run("should work without bounces", func(t *testing.T) {
		gotPosition, gotBounces := CalculateNextPosition(1, 3, 2)
		utils.AssertEquals(t, gotPosition, 3)
		utils.AssertEquals(t, gotBounces, 0)
	})
	t.Run("should work backwards without bounces", func(t *testing.T) {
		gotPosition, gotBounces := CalculateNextPosition(3, 3, -3)
		utils.AssertEquals(t, gotPosition, 0)
		utils.AssertEquals(t, gotBounces, 0)
	})
	t.Run("should work with bounces", func(t *testing.T) {
		gotPosition, gotBounces := CalculateNextPosition(1, 3, 8)
		utils.AssertEquals(t, gotPosition, 3)
		utils.AssertEquals(t, gotBounces, 2)
	})
}
