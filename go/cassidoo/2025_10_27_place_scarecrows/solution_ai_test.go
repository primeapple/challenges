package placescarecrows

import (
	"cassidoo/utils"
	"testing"
)

func TestPlaceScarecrowsAI(t *testing.T) {
	// OWN STUFF
	t.Run("with k=1 should cover i", func(t *testing.T) {
		field := []int{1, 1, 1}
		k := 1

		got := PlaceScarecrowsAI(field, k)
		want := []int{0, 1, 2}

		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("with k=3 should cover i-1, i, i+1", func(t *testing.T) {
		field := []int{1, 1, 1}
		k := 3

		got := PlaceScarecrowsAI(field, k)
		want := []int{1}

		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("place scarecrows on unneeded fields if it makes sense", func(t *testing.T) {
		field := []int{1, 0, 1}
		k := 3

		got := PlaceScarecrowsAI(field, k)
		want := []int{1}

		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("ignore protection for unneeded fields", func(t *testing.T) {
		field := []int{0, 1, 1}
		k := 1

		got := PlaceScarecrowsAI(field, k)
		want := []int{1, 2}

		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("place scarecrow on the later field if we have the choice", func(t *testing.T) {
		field := []int{0, 1, 1}
		k := 3

		got := PlaceScarecrowsAI(field, k)
		want := []int{2}

		utils.AssertDeepEquals(t, got, want)
	})

	// EXAMPLES
	t.Run("should work with example 1", func(t *testing.T) {
		field := []int{1, 1, 0, 1, 1, 0, 1}
		k := 3

		got := PlaceScarecrowsAI(field, k)
		want := []int{1, 4, 6}

		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		field := []int{1, 0, 1, 1, 0, 1}
		k := 3

		got := PlaceScarecrowsAI(field, k)
		want := []int{1, 4}

		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 3", func(t *testing.T) {
		field := []int{1, 1, 1, 1, 1}
		k := 1

		got := PlaceScarecrowsAI(field, k)
		want := []int{0, 1, 2, 3, 4}

		utils.AssertDeepEquals(t, got, want)
	})
}
