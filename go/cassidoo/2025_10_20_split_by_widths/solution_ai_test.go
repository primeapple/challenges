package splitbywidths

import (
	"reflect"
	"testing"
)

func TestSplitByWidthsAI(t *testing.T) {
	t.Run("should work with fitting string to widths", func(t *testing.T) {
		str := "abcd"
		widths := []int{1, 3}

		got := SplitByWidthsAI(str, widths)
		want := []string{"a", "bcd"}

		assertDeepEqualsAI(t, got, want)
	})

	t.Run("should work with shorter string than widths", func(t *testing.T) {
		str := "abcd"
		widths := []int{1, 4}

		got := SplitByWidthsAI(str, widths)
		want := []string{"a", "bcd"}

		assertDeepEqualsAI(t, got, want)
	})

	t.Run("should work with longer string than widths", func(t *testing.T) {
		str := "abcd"
		widths := []int{1, 2}

		got := SplitByWidthsAI(str, widths)
		want := []string{"a", "bc", "d"}

		assertDeepEqualsAI(t, got, want)
	})

	t.Run("should work with example", func(t *testing.T) {
		str := "Supercalifragilisticexpialidocious"
		widths := []int{5, 9, 4}

		got := SplitByWidthsAI(str, widths)
		want := []string{"Super", "califragi", "list", "icex", "pial", "idoc", "ious"}

		assertDeepEqualsAI(t, got, want)
	})

	t.Run("should work with empty string", func(t *testing.T) {
		str := ""
		widths := []int{1, 2}

		got := SplitByWidthsAI(str, widths)
		want := []string{}

		assertDeepEqualsAI(t, got, want)
	})

	t.Run("should work with empty widths slice", func(t *testing.T) {
		str := "abcd"
		widths := []int{}

		got := SplitByWidthsAI(str, widths)
		want := []string{}

		assertDeepEqualsAI(t, got, want)
	})
}

func assertDeepEqualsAI(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v expected %v", got, want)
	}
}
