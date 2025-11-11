package merge

func Merge(mergeInto, merger []int) []int {
	positionMergeInto := 0
	positionMerger := 0

	for positionMerger < len(merger) {
		currentMergeIntoElement := mergeInto[positionMergeInto]
		currentMergerElement := merger[positionMerger]

		if currentMergeIntoElement == 0 || currentMergerElement < currentMergeIntoElement {
			copy(mergeInto[positionMergeInto+1:], mergeInto[positionMergeInto:])
			mergeInto[positionMergeInto] = currentMergerElement
			positionMerger++
		} else {
			positionMergeInto++
		}
	}
	return mergeInto
}
