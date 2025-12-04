package merge

import (
	"cassidoo/utils"
	"testing"
)

func TestMerge(t *testing.T) {
	// EXAMPLE
	t.Run("should work with example", func(t *testing.T) {
		mergeInto := []int{1, 3, 5, 0, 0, 0}
		merger := []int{2, 4, 6}

		got := Merge(mergeInto, merger)
		want := []int{1, 2, 3, 4, 5, 6}

		utils.AssertDeepEquals(t, got, want)
	})
}
