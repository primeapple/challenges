package placescarecrows

func PlaceScarecrowsAI(field []int, k int) []int {
	n := len(field)
	if n == 0 {
		return []int{}
	}

	// Calculate the radius (how far from center the scarecrow protects)
	radius := k / 2

	result := []int{}
	protected := make([]bool, n)

	for i := 0; i < n; i++ {
		// If this position needs protection and isn't already protected
		if field[i] == 1 && !protected[i] {
			// Place scarecrow as far right as possible while still covering position i
			scarecrowPos := min(i+radius, n-1)

			result = append(result, scarecrowPos)

			// Mark all positions this scarecrow protects
			start := max(0, scarecrowPos-radius)
			end := min(n-1, scarecrowPos+radius)
			for j := start; j <= end; j++ {
				protected[j] = true
			}
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
