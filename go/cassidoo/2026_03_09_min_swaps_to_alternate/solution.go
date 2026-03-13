package minSwapsToAlternate

func MinSwapsToAlternate(pattern string) int {
	minSwapsStartingWithB := calcMinSwaps([]rune(pattern), 'b', 'a')
	minSwapsStartingWithA := calcMinSwaps([]rune(pattern), 'a', 'b')

	return min(minSwapsStartingWithA, minSwapsStartingWithB)
}

func calcMinSwaps(elements []rune, oddChar, evenChar rune) int {
	swaps := 0
	for i, char := range elements {
		switch char {
		case oddChar:
			if i%2 == 1 {
				continue
			}
		case evenChar:
			if i%2 == 0 {
				continue
			}
		}


		indexToSwap := findNextIndexWithDifferentChar(elements, i)
		if indexToSwap == -1 {
			return -1
		}

		elements[i], elements[indexToSwap] = elements[indexToSwap], elements[i]
		swaps = swaps + indexToSwap - i
	}

	return swaps
}

func findNextIndexWithDifferentChar(elements []rune, currentIndex int) int {
	for i:=currentIndex+1; i<len(elements); i++ {
		if elements[i] != elements[currentIndex] {
			return i
		}
	}
	return -1
}
