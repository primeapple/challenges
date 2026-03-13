package minSwapsToAlternate

import (
	"cassidoo/utils"
	"testing"
)

func TestMinSwapsToAlternateAI(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := MinSwapsToAlternateAI("aabb")
		want := 1
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := MinSwapsToAlternateAI("aaab")
		want := -1
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := MinSwapsToAlternateAI("aaaabbbb")
		want := 6
		utils.AssertEquals(t, got, want)
	})
}
