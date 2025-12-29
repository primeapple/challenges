package latinSquare

import (
	"cassidoo/utils"
	"testing"
)

func TestLatinSquare(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := LatinSquare(1)
		want := [][]int{{1}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := LatinSquare(2)
		want := [][]int{{1, 2}, {2, 1}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := LatinSquare(4)
		want := [][]int{{1, 2, 3, 4}, {2, 3, 4, 1}, {3, 4, 1, 2}, {4, 1, 2, 3}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should give correct length on very long example", func(t *testing.T) {
		got := len(LatinSquare(10000))
		want := 10000
		utils.AssertEquals(t, got, want)
	})
}

func TestLatinSquareConcurrent(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := LatinSquareConcurrent(1)
		want := [][]int{{1}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := LatinSquareConcurrent(2)
		want := [][]int{{1, 2}, {2, 1}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := LatinSquareConcurrent(4)
		want := [][]int{{1, 2, 3, 4}, {2, 3, 4, 1}, {3, 4, 1, 2}, {4, 1, 2, 3}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should give correct length on very long example", func(t *testing.T) {
		got := len(LatinSquareConcurrent(10000))
		want := 10000
		utils.AssertEquals(t, got, want)
	})
}
