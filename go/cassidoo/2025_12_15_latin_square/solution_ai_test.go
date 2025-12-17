package latinSquare

import (
	"cassidoo/utils"
	"testing"
)

func TestLatinSquareAI(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := LatinSquareAI(1)
		want := [][]int{{1}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := LatinSquareAI(2)
		want := [][]int{{1, 2}, {2, 1}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := LatinSquareAI(4)
		want := [][]int{{1, 2, 3, 4}, {2, 3, 4, 1}, {3, 4, 1, 2}, {4, 1, 2, 3}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should give correct length on very long example", func(t *testing.T) {
		got := len(LatinSquareAI(10000))
		want := 10000
		utils.AssertEquals(t, got, want)
	})
}

func TestLatinSquareConcurrentAI(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := LatinSquareConcurrentAI(1)
		want := [][]int{{1}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := LatinSquareConcurrentAI(2)
		want := [][]int{{1, 2}, {2, 1}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := LatinSquareConcurrentAI(4)
		want := [][]int{{1, 2, 3, 4}, {2, 3, 4, 1}, {3, 4, 1, 2}, {4, 1, 2, 3}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should give correct length on very long example", func(t *testing.T) {
		got := len(LatinSquareConcurrentAI(10000))
		want := 10000
		utils.AssertEquals(t, got, want)
	})
}
