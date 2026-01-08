package maxScoreWithOneReset

// MaxScoreWithOneResetAI calculates the maximum score by summing array elements
// with the option to reset the score once to 0 and continue.
// Uses dynamic programming to track current sum with and without using the reset.
func MaxScoreWithOneResetAI(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Track two states at each position:
	// sumNoReset: current running sum without ever using the reset
	// sumWithReset: current running sum having used the reset at some point
	sumNoReset := 0
	sumWithReset := 0
	maxResult := 0

	for _, num := range nums {
		// Update sumWithReset: we can either
		// 1. Continue with previous sumWithReset (already reset before) + num
		// 2. Reset right now (take previous sumNoReset, reset to 0, then add num)
		newSumWithReset := max(sumWithReset+num, num)

		// Update sumNoReset: just add the current number
		sumNoReset = sumNoReset + num

		sumWithReset = newSumWithReset

		// Track the maximum score we can achieve at any point
		// We can stop at any position
		maxResult = max(maxResult, max(sumNoReset, sumWithReset))
	}

	return maxResult
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
