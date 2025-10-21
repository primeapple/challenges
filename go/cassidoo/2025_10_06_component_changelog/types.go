package groupChangelogEdits

type Edit struct {
	component string
	timestamp date
}

type GroupedEdits struct {
	component string
	start     date
	end       date
}
