package repeatedIntegers

func RepeatedIntegers(n int) []int {
	result := []int{}

	for number := 1; number <= n; number++ {
		for repeat := 0; repeat < number; repeat++ {
			result = append(result, number)
		}
	}

	return result
}
