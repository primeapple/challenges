package groupChangelogEdits

import (
	"cassidoo/utils"
	"fmt"
	"slices"
	"time"
)

const GROUP_WINDOW_DURATION = 10 * time.Minute

func GroupChangelogEdits(edits []Edit) []GroupedEdit {
	slices.SortFunc(edits, func(e1, e2 Edit) int {
		return e1.timestamp.Compare(e2.timestamp)
	})

	return utils.Reduce(edits, make([]GroupedEdit, 0), func(edit Edit, groups []GroupedEdit) []GroupedEdit {
		latestComponentGroup, err := findLatestComponentGroup(groups, edit.component)
		if err != nil {
			return append(groups, GroupedEdit{component: edit.component, start: edit.timestamp, end: edit.timestamp})
		}

		if edit.timestamp.Sub(latestComponentGroup.start) <= GROUP_WINDOW_DURATION {
			latestComponentGroup.end = edit.timestamp
		} else {
			return append(groups, GroupedEdit{component: edit.component, start: edit.timestamp, end: edit.timestamp})
		}

		return groups
	})
}

func findLatestComponentGroup(edits []GroupedEdit, component string) (*GroupedEdit, error) {
	for i := len(edits) - 1; i >= 0; i-- {
		if edits[i].component == component {
			return &edits[i], nil
		}
	}

	return nil, fmt.Errorf("No group found for component %q", component)
}
