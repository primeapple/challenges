package minRepairs

import (
	"cassidoo/utils"
	"testing"
)

func TestMinRepairs(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := MinRepairs([][]int{
			{1, 0, 0, 1},
			{1, 0, 0, 1},
			{1, 1, 0, 1},
			{0, 1, 1, 1},
		}, 2)
		want := 2
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := MinRepairs([][]int{
			{1, 0, 0, 1},
			{1, 0, 0, 1},
			{1, 1, 0, 1},
			{0, 0, 1, 1},
		}, 1)
		want := 3
		utils.AssertEquals(t, got, want)
	})
}

func TestFindBrokenComponents(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := FindBrokenComponents([][]int{
			{1, 0, 0, 1},
			{1, 0, 0, 1},
			{1, 1, 0, 1},
			{0, 1, 1, 1},
		})
		want := []Component{
			{{0, 1}, {0, 2}, {1, 1}, {1, 2}, {2,2}},
			{{3, 0}},
		}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := FindBrokenComponents([][]int{
			{1, 0, 0, 1},
			{1, 0, 0, 1},
			{1, 1, 0, 1},
			{0, 0, 1, 1},
		})
		want := []Component{
			{{0, 1}, {0, 2}, {1, 1}, {1, 2}, {2, 2}},
			{{3, 0}, {3, 1}},
		}
		utils.AssertDeepEquals(t, got, want)
	})
}
