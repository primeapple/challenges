package maxSubarraySum

import (
	"cassidoo/utils"
	"testing"
)

func TestMaxSubarraySum(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := MaxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
		want := 6
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := MaxSubarraySum([]int{5})
		want := 5
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := MaxSubarraySum([]int{-1, -2, -3, -4})
		want := -1
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 4", func(t *testing.T) {
		got := MaxSubarraySum([]int{5, 4, -1, 7, 8})
		want := 23
		utils.AssertEquals(t, got, want)
	})
}
