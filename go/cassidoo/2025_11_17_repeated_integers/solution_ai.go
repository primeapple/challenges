package repeatedIntegers

func RepeatedIntegersAI(n int) []int {
	// Calculate the total size needed: sum of 1 to n = n*(n+1)/2
	size := n * (n + 1) / 2
	result := make([]int, 0, size)

	// For each integer i from 1 to n, append it i times
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			result = append(result, i)
		}
	}

	return result
}
