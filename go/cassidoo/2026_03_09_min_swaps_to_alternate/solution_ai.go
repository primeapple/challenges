package minSwapsToAlternate

func MinSwapsToAlternateAI(s string) int {
	countA := 0
	countB := 0
	for _, c := range s {
		if c == 'a' {
			countA++
		} else {
			countB++
		}
	}

	diffCounts := countA - countB
	if diffCounts < 0 {
		diffCounts = -diffCounts
	}
	if diffCounts > 1 {
		return -1
	}

	aPositions := make([]int, 0, countA)
	for i, c := range s {
		if c == 'a' {
			aPositions = append(aPositions, i)
		}
	}

	swapsStartingWithA := 0
	for k := 0; k < countA; k++ {
		diff := aPositions[k] - 2*k
		if diff < 0 {
			diff = -diff
		}
		swapsStartingWithA += diff
	}

	swapsStartingWithB := 0
	for k := 0; k < countA; k++ {
		diff := aPositions[k] - (2*k + 1)
		if diff < 0 {
			diff = -diff
		}
		swapsStartingWithB += diff
	}

	if countA > countB {
		return swapsStartingWithA
	} else if countB > countA {
		return swapsStartingWithB
	}

	if swapsStartingWithA < swapsStartingWithB {
		return swapsStartingWithA
	}
	return swapsStartingWithB
}
