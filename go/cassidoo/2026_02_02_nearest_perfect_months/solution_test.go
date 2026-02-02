package nearestPerfectMonths

import (
	"cassidoo/utils"
	"testing"
)

func TestNearestPerfectMonths(t *testing.T) {
	t.Run("should work with example 1", func(t *testing.T) {
		got := NearestPerfectMonths(2025)
		want := Solution{prev: "2021-02", next: "2026-02"}
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should work with example 2", func(t *testing.T) {
		got := NearestPerfectMonths(2026)
		want := Solution{prev: "2026-02", next: "2027-02"}
		utils.AssertDeepEquals(t, got, want)
	})
}
