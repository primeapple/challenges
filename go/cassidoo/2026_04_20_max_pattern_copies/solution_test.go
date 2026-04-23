package maxPatternCopies

import (
	"cassidoo/utils"
	"testing"
)

func TestMaxPatternCopies(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := MaxPatternCopies("abcabc???", "ac")
		want := 3
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := MaxPatternCopies("aab??", "aab")
		want := 1
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := MaxPatternCopies("??????", "abc")
		want := 2
		utils.AssertEquals(t, got, want)
	})
}
