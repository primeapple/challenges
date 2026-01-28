package groupChangelogEdits


import (
	"sort"
	"time"
)

func GroupChangelogEditsAI(edits []Edit) []GroupedEdit {
	// Sort edits by timestamp
	sort.Slice(edits, func(i, j int) bool {
		return edits[i].timestamp.Before(edits[j].timestamp)
	})

	// Group timestamps by component
	componentEdits := make(map[string][]time.Time)
	for _, edit := range edits {
		componentEdits[edit.component] = append(componentEdits[edit.component], edit.timestamp)
	}

	var result []GroupedEdit
	for comp, ts := range componentEdits {
		if len(ts) == 0 {
			continue
		}
		// Sort timestamps for this component (though already sorted globally)
		sort.Slice(ts, func(i, j int) bool {
			return ts[i].Before(ts[j])
		})

		// Group consecutive timestamps within 10 minutes
		start := ts[0]
		end := ts[0]
		for i := 1; i < len(ts); i++ {
			if ts[i].Sub(end) <= 10*time.Minute {
				end = ts[i]
			} else {
				result = append(result, GroupedEdit{
					component: comp,
					start:     start,
					end:       end,
				})
				start = ts[i]
				end = ts[i]
			}
		}
		result = append(result, GroupedEdit{
			component: comp,
			start:     start,
			end:       end,
		})
	}

	// Sort result by start time
	sort.Slice(result, func(i, j int) bool {
		return result[i].start.Before(result[j].start)
	})

	return result
}
