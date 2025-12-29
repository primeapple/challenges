package replaceRepeats

import (
	"cassidoo/utils"
	"testing"
)

func TestReplaceRepeatsAI(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got, error := ReplaceRepeatsAI("1234500362000440", 0)
		want := "1234523623441"
		utils.AssertEquals(t, got, want)
		utils.AssertNil(t, error)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got, error := ReplaceRepeatsAI("000000000000", 0)
		want := "12"
		utils.AssertEquals(t, got, want)
		utils.AssertNil(t, error)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		got, error := ReplaceRepeatsAI("123456789", 1)
		want := "123456789"
		utils.AssertEquals(t, got, want)
		utils.AssertNil(t, error)
	})

	t.Run("should throw an error on non numbers", func(t *testing.T) {
		_, err := ReplaceRepeatsAI("1a", 1)
		utils.AssertError(t, err, ErrInvalidDigit)
	})
}
