package minassemblytime

import (
	"slices"
)

func MinAssemblyTime(parts []AssemblyDetail) int {
	currentTime := 0

	sortedParts := slices.Clone(parts)
	slices.SortFunc(sortedParts, func(p1, p2 AssemblyDetail) int {
		if p1.arrivalDays < p2.arrivalDays {
			return -1
		} else {
			return 1
		}
	})

	for _, part := range sortedParts {
		// the part is not there yet, we have to wait
		if currentTime < part.arrivalDays*24 {
			currentTime = part.arrivalDays * 24
		}

		currentTime += part.assemblyHours
	}

	return currentTime
}
