package perrinCombinations

import (
	"cassidoo/utils"
	"testing"
)

func TestPerrinCombinations(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := PerrinCombinations(8, 12)
		want := [][]int{{0, 2, 3, 7}, {0, 5, 7}, {2, 3, 7}, {5, 7}}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := PerrinCombinations(6, 5)
		want := [][]int{{0, 2, 3}, {0, 5}, {2, 3}, {5}}
		utils.AssertDeepEquals(t, got, want)
	})
}
