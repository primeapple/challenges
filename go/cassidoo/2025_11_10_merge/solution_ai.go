package merge

func MergeAI(mergeInto, merger []int) []int {
	// Find last non-zero position in mergeInto
	lastMergeInto := 0
	for i := 0; i < len(mergeInto); i++ {
		if mergeInto[i] == 0 {
			break
		}
		lastMergeInto = i
	}

	// Merge from back to front
	positionMergeInto := lastMergeInto
	positionMerger := len(merger) - 1
	positionResult := len(mergeInto) - 1

	for positionMerger >= 0 {
		if positionMergeInto >= 0 && mergeInto[positionMergeInto] > merger[positionMerger] {
			mergeInto[positionResult] = mergeInto[positionMergeInto]
			positionMergeInto--
		} else {
			mergeInto[positionResult] = merger[positionMerger]
			positionMerger--
		}
		positionResult--
	}

	return mergeInto
}
