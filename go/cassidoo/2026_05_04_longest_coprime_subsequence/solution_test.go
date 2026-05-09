package longestCoprimeSubsequence

import (
	"cassidoo/utils"
	"testing"
)

func TestLongestCoprimeSubsequence(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := LongestCoprimeSubsequence([]int{6, 12, 4, 8})
		want := 1
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := LongestCoprimeSubsequence([]int{4, 3, 6, 9, 7, 2})
		want := 4
		utils.AssertEquals(t, got, want)
	})
}

func TestSubsequences(t *testing.T) {
	t.Run("should work with slices of length 1", func(t *testing.T) {
		got := Subsequences([]int{1})
		want := [][]int{{1}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with slices of length 2", func(t *testing.T) {
		got := Subsequences([]int{1, 2})
		want := [][]int{{1, 2}, {2}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with slices of length 3", func(t *testing.T) {
		got := Subsequences([]int{3, 2, 1})
		want := [][]int{{3, 2, 1}, {2, 1}, {3, 1}, {1}}
		utils.AssertDeepEquals(t, got, want)
	})
}
