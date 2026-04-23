package maxPatternCopies

func MaxPatternCopiesAI(s, pattern string) int {
	sCount := map[rune]int{}
	for _, r := range s {
		sCount[r]++
	}

	pCount := map[rune]int{}
	for _, r := range pattern {
		pCount[r]++
	}

	wildcards := sCount['?']
	lo, hi := 0, len(s)/len(pattern)

	for lo < hi {
		mid := (lo + hi + 1) / 2
		needed := 0
		for c, count := range pCount {
			if deficit := count*mid - sCount[c]; deficit > 0 {
				needed += deficit
			}
		}
		if needed <= wildcards {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	return lo
}
