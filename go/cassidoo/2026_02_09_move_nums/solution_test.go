package moveNums

import (
	"cassidoo/utils"
	"testing"
)

func TestMoveNums(t *testing.T) {
	t.Run("should work with example", func(t *testing.T) {
		got := MoveNums([]int{0, 2, 0, 3, 10}, 0)
		want := []int{2, 3, 10, 0, 0}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work without n appearing in nums", func(t *testing.T) {
		got := MoveNums([]int{1, 2, 3, 4}, 0)
		want := []int{1, 2, 3, 4}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with last values being", func(t *testing.T) {
		got := MoveNums([]int{1, 2, 3, 3}, 3)
		want := []int{1, 2, 3, 3}
		utils.AssertDeepEquals(t, got, want)
	})
}
