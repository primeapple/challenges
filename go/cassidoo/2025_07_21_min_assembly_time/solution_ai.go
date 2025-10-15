package minassemblytime

import "sort"

func MinAssemblyTimeAI(parts []AssemblyDetail) int {
	if len(parts) == 0 {
		return 0
	}

	// Create a copy of parts to avoid modifying the original
	sortedParts := make([]AssemblyDetail, len(parts))
	copy(sortedParts, parts)

	// Sort parts by arrival time (earliest first)
	sort.Slice(sortedParts, func(i, j int) bool {
		return sortedParts[i].arrivalDays < sortedParts[j].arrivalDays
	})

	currentTime := 0

	for _, part := range sortedParts {
		// Convert arrival days to hours
		arrivalTime := part.arrivalDays * 24

		// We can start working on this part at the maximum of:
		// - When we finish the previous part (currentTime)
		// - When this part actually arrives (arrivalTime)
		startTime := max(currentTime, arrivalTime)

		// Update current time to when we finish this part
		currentTime = startTime + part.assemblyHours
	}

	return currentTime
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}