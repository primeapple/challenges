package majorityElement

import (
	"cassidoo/utils"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got, _ := MajorityElement([]int{2, 2, 1, 1, 2, 2, 1, 2, 2})
		want := 2
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got, _ := MajorityElement([]int{3, 3, 4, 2, 3, 3, 1})
		want := 3
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should give an error if no majority for a single element", func(t *testing.T) {
		_, got := MajorityElement([]int{1,1,2,2})
		utils.AssertError(t, got, ErrNoMajorityFound)
	})
}
