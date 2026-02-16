package whatKindOfNumber

import (
	"cassidoo/utils"
	"testing"
)

func TestWhatKindOfNumber(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := WhatKindOfNumber(6)
		want := Perfect
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := WhatKindOfNumber(12)
		want := Abundant
		utils.AssertEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got := WhatKindOfNumber(4)
		want := Deficient
		utils.AssertEquals(t, got, want)
	})
}
