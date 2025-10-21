package groupChangelogEdits

import (
	"reflect"
	"testing"
)

func TestGroupChangelogEdits(t *testing.T) {
	t.Run("should work with example", func(t *testing.T) {
		edits := []Edit{
			{timestamp: "2025-10-06T08:00:00Z", component: "Header"},
			{timestamp: "2025-10-06T08:05:00Z", component: "Header"},
			{timestamp: "2025-10-06T08:20:00Z", component: "Header"},
			{timestamp: "2025-10-06T08:07:00Z", component: "Footer"},
			{timestamp: "2025-10-06T08:15:00Z", component: "Footer"},
		}

		got := GroupChangelogEdits(edits)
		want := []GroupedEdits{
			{component: "Header", start: "2025-10-06T08:00:00Z", end: "2025-10-06T08:05:00Z"},
			{component: "Header", start: "2025-10-06T08:20:00Z", end: "2025-10-06T08:20:00Z"},
			{component: "Footer", start: "2025-10-06T08:07:00Z", end: "2025-10-06T08:15:00Z"},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v expected %v", got, want)
		}
	})
}
