package sortHtmlColors

import (
	"testing"
)

func TestSortHtmlColorsAI(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("should work with permutation", func(t *testing.T) {
			got := SortHtmlColorsAI(GetColorsInRandomOrder())
			want := HtmlColors
			assertDeepEquals(t, got, want)
		})
	}
}
