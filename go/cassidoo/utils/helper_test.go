package utils

import "testing"

func TestPop(t *testing.T) {
	t.Run("should give first element and remove it", func(t *testing.T) {
		list := []int{1, 2}
		got, ok := Pop(&list)
		want := 1
		AssertTrue(t, ok)
		AssertEquals(t, got, want)
		AssertDeepEquals(t, list, []int{2})
	})

	t.Run("should return zero value if list is empty", func(t *testing.T) {
		list := []int{}
		got, ok := Pop(&list)
		want := 0
		AssertFalse(t, ok)
		AssertEquals(t, got, want)
		AssertDeepEquals(t, list, []int{})
	})
}
