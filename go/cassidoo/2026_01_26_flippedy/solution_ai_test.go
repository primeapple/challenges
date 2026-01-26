package flippedy

import (
	"cassidoo/utils"
	"testing"
)

func TestFlippedyAI(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := FlippedyAI("cat and mice")
		want := "cat dna mice"
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := FlippedyAI("banana healthy")
		want := "banana healthy"
		utils.AssertEquals(t, got, want)
	})
}
