package hyperfactorial

import (
	"cassidoo/utils"
	"testing"
)

func TestHyperfactorial(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := Hyperfactorial(0)
		want := 1
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := Hyperfactorial(2)
		want := 4
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := Hyperfactorial(3)
		want := 108
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 4", func(t *testing.T) {
		got := Hyperfactorial(7)
		want := 3319766398771200000
		utils.AssertEquals(t, got, want)
	})
}
