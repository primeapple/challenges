package repeatedIntegers

import (
	"testing"
)

func TestRepeatedIntegers(t *testing.T) {
	// EXAMPLE
	t.Run("should work with example", func(t *testing.T) {
		n := 4

		got := RepeatedIntegers(n)
		want := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}

		assertDeepEquals(t, got, want)
	})
}
