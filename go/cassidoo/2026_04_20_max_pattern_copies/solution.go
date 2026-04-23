package maxPatternCopies

func MaxPatternCopies(s, pattern string) int {
	occurenceMap := map[rune]int{}
	for _, r := range s {
		currentCount, ok := occurenceMap[r]
		if !ok {
			occurenceMap[r] = 1
		} else {
			occurenceMap[r] = currentCount + 1
		}
	}

	patternCount := 0
	for {
		ok := updateOccurenceMap(pattern, occurenceMap)
		if ok {
			patternCount++
		} else {
			break
		}
	}

	return patternCount
}

func updateOccurenceMap(pattern string, occurenceMap map[rune]int) bool {
	for _, r := range pattern {
		currentCount, ok := occurenceMap[r]
		if !ok || currentCount == 0 {
			placeholderCount, ok := occurenceMap['?']
			if !ok || placeholderCount == 0 {
				return false
			} else {
				occurenceMap['?'] = placeholderCount - 1
			}
		} else {
			occurenceMap[r] = currentCount - 1
		}
	}

	return true
}
