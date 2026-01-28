package groupChangelogEdits

import "time"

type Edit struct {
	component string
	timestamp time.Time
}

type GroupedEdit struct {
	component string
	end       time.Time
	start     time.Time
}
