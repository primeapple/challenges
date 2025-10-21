package splitbywidths

import (
	"reflect"
	"testing"
)

func TestSplitByWidths(t *testing.T) {
	t.Run("should work with fitting string to widths", func(t *testing.T) {
		str := "abcd"
		widths := []int{1, 3}

		got := SplitByWidths(str, widths)
		want := []string{"a", "bcd"}

		assertDeepEquals(t, got, want)
	})

	t.Run("should work with shorter string than widths", func(t *testing.T) {
		str := "abcd"
		widths := []int{1, 4}

		got := SplitByWidths(str, widths)
		want := []string{"a", "bcd"}

		assertDeepEquals(t, got, want)
	})

	t.Run("should work with longer string than widths", func(t *testing.T) {
		str := "abcd"
		widths := []int{1, 2}

		got := SplitByWidths(str, widths)
		want := []string{"a", "bc", "d"}

		assertDeepEquals(t, got, want)
	})

	t.Run("should work with example", func(t *testing.T) {
		str := "Supercalifragilisticexpialidocious"
		widths := []int{5, 9, 4}

		got := SplitByWidths(str, widths)
		want := []string{"Super", "califragi", "list", "icex", "pial", "idoc", "ious"}

		assertDeepEquals(t, got, want)
	})

	t.Run("should work with empty string", func(t *testing.T) {
		str := ""
		widths := []int{1, 2}

		got := SplitByWidths(str, widths)
		want := []string{}

		assertDeepEquals(t, got, want)
	})

	t.Run("should work with empty widths slice", func(t *testing.T) {
		str := "abcd"
		widths := []int{}

		got := SplitByWidths(str, widths)
		want := []string{}

		assertDeepEquals(t, got, want)
	})
}

func assertDeepEquals(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v expected %v", got, want)
	}
}
