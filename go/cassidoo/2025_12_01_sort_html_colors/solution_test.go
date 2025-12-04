package sortHtmlColors

import (
	"testing"
)

func TestSortHtmlColors(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("should work with permutation", func(t *testing.T) {
			got := SortHtmlColors(GetColorsInRandomOrder())
			want := HtmlColors
			assertDeepEquals(t, got, want)
		})
	}
}

func TestCompareString(t *testing.T) {
	t.Run("returns positive number if first string is bigger", func(t *testing.T) {
		got := CompareString("b", "a")
		assertIsPositive(t, got)
	})

	t.Run("returns negative number if first string is smaller", func(t *testing.T) {
		got := CompareString("a", "b")
		assertIsNegative(t, got)
	})

	t.Run("returns positive number if first string has the second one as prefix", func(t *testing.T) {
		got := CompareString("aa", "a")
		assertIsPositive(t, got)
	})

	t.Run("returns negative number if second string has the first one as prefix", func(t *testing.T) {
		got := CompareString("a", "aa")
		assertIsNegative(t, got)
	})
}
