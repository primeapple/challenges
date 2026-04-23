package maxPatternCopies

import (
	"cassidoo/utils"
	"testing"
)

func TestMaxPatternCopiesAI(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := MaxPatternCopiesAI("abcabc???", "ac")
		want := 3
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := MaxPatternCopiesAI("aab??", "aab")
		want := 1
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := MaxPatternCopiesAI("??????", "abc")
		want := 2
		utils.AssertEquals(t, got, want)
	})
}
