package maxScoreWithOneReset

func MaxScoreWithOneReset(numbers []int) int {
	maxScore := 0
	current := 0
	for i := len(numbers) - 1; i >= 0; i-- {
		number := numbers[i]

		current = current + number

		if (current > maxScore) {
			maxScore = current
		}
	}
	return maxScore
}
