package alternatingArray

import (
	"cassidoo/utils"
	"testing"
)

func TestAlternatingArray(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		array := []int{}
		got := AlternatingArray(array)
		utils.AssertTrue(t, got)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		array := []int{1}
		got := AlternatingArray(array)
		utils.AssertTrue(t, got)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		array := []int{1, 1}
		got := AlternatingArray(array)
		utils.AssertTrue(t, got)
	})

	t.Run("should work with example 4", func(t *testing.T) {
		array := []int{1, 2, 1}
		got := AlternatingArray(array)
		utils.AssertTrue(t, got)
	})

	t.Run("should work with example 5", func(t *testing.T) {
		array := []int{10, 5, 10, 5, 10}
		got := AlternatingArray(array)
		utils.AssertTrue(t, got)
	})

	t.Run("should work with example 6", func(t *testing.T) {
		array := []int{2, 2, 3, 3}
		got := AlternatingArray(array)
		utils.AssertFalse(t, got)
	})

	t.Run("should work with example 1", func(t *testing.T) {
		array := []int{5, 4, 3, 5, 4, 3}
		got := AlternatingArray(array)
		utils.AssertFalse(t, got)
	})

}
