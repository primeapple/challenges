package maxSubarraySum

func MaxSubarraySum(numbers []int) int {
	best := numbers[0]
	currentSum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		number := numbers[i]

		currentSum = max(number, currentSum+number)
		best = max(best, currentSum)
	}
	return best
}
