package fuzzySearch

import (
	"cassidoo/utils"
	"testing"
)

func TestFuzzySearch(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := FuzzySearch("the cat sat on the mat", "cat", 0)
		want := []Result{{position: 4, errors: 0}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := FuzzySearch("cassidoo", "cool", 1)
		want := []Result{}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := FuzzySearch("cassidoo", "cool", 3)
		want := []Result{{position: 0, errors: 3}, {position: 4, errors: 3}, {position: 5, errors: 2}, {position: 6, errors: 2}, {position: 7, errors: 3}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("deletion match when text is shorter than pattern", func(t *testing.T) {
		got := FuzzySearch("ab", "abcd", 2)
		want := []Result{{position: 0, errors: 2}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("deletion match position is correct", func(t *testing.T) {
		got := FuzzySearch("xac", "abc", 1)
		want := []Result{{position: 1, errors: 1}}
		utils.AssertDeepEquals(t, got, want)
	})
}

func TestExactMatch(t *testing.T) {
	t.Run("should return position of exact match", func(t *testing.T) {
		got := ExactMatch("the cat sat on the mat", "cat")
		want := 4
		utils.AssertEquals(t, got, want)
	})

	t.Run("should return -1 if no exact match", func(t *testing.T) {
		got := ExactMatch("the rat sat on the mat", "cat")
		want := -1
		utils.AssertEquals(t, got, want)
	})
}
