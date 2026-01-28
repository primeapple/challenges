package groupChangelogEdits

import (
	"cassidoo/utils"
	"testing"
	"time"
)

func TestGroupChangelogEdits(t *testing.T) {
	t.Run("should work with example", func(t *testing.T) {
		got := GroupChangelogEdits([]Edit{
			{timestamp: time.Date(2025, time.October, 6, 8, 0, 0, 0, time.UTC), component: "Header"},
			{timestamp: time.Date(2025, time.October, 6, 8, 5, 0, 0, time.UTC), component: "Header"},
			{timestamp: time.Date(2025, time.October, 6, 8, 20, 0, 0, time.UTC), component: "Header"},
			{timestamp: time.Date(2025, time.October, 6, 8, 7, 0, 0, time.UTC), component: "Footer"},
			{timestamp: time.Date(2025, time.October, 6, 8, 15, 0, 0, time.UTC), component: "Footer"},
		})
		want := []GroupedEdit{
			{
				component: "Header",
				start:     time.Date(2025, time.October, 6, 8, 0, 0, 0, time.UTC),
				end:       time.Date(2025, time.October, 6, 8, 5, 0, 0, time.UTC),
			},
			{
				component: "Footer",
				start:     time.Date(2025, time.October, 6, 8, 7, 0, 0, time.UTC),
				end:       time.Date(2025, time.October, 6, 8, 15, 0, 0, time.UTC),
			},
			{
				component: "Header",
				start:     time.Date(2025, time.October, 6, 8, 20, 0, 0, time.UTC),
				end:       time.Date(2025, time.October, 6, 8, 20, 0, 0, time.UTC),
			},
		}

		utils.AssertEquals(t, len(got), len(want))
		utils.AssertDeepEquals(t, got, want)
	})
}
