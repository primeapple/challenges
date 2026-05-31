package shuffleLine

import (
	"cassidoo/utils"
	"testing"
)

func TestShuffleLine(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := ShuffleLine([]string{"Ada", "Ben", "Cam", "Diya", "Eli", "Fay"}, 3)
		want := []string{"Ada", "Ben", "Diya", "Eli", "Cam", "Fay"}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := ShuffleLine([]string{"A", "B", "C", "D", "E"}, 2)
		want := []string{"A", "C", "E", "B", "D"}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := ShuffleLine([]string{"Mo", "Noah", "Oli"}, 1)
		want := []string{"Mo", "Noah", "Oli"}
		utils.AssertDeepEquals(t, got, want)
	})
}
