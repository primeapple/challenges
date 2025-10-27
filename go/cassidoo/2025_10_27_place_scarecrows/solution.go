package placescarecrows

func PlaceScarecrows(field []int, k int) []int {
	placements := []int{}

	coverageMiddle := k / 2
	coveredFieldsFromNow := 0

	for index, position := range field {
		if needsProtection(position) && coveredFieldsFromNow == 0 {
			placements = append(placements, min(index+coverageMiddle, len(field)-1))
			coveredFieldsFromNow = k
		}

		coveredFieldsFromNow = max(0, coveredFieldsFromNow-1)
	}

	return placements
}

func needsProtection(position int) bool {
	if position == 1 {
		return true
	}

	return false
}
