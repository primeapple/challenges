package whatKindOfNumber

// WhatKindOfNumberAI determines if a number is abundant, deficient, or perfect
// based on the sum of its proper divisors (divisors less than the number itself)
func WhatKindOfNumberAI(n int) KindOfNumber {
	// Calculate sum of proper divisors (all divisors except n itself)
	divisorSum := sumOfProperDivisors(n)

	// Compare sum to the number to classify it
	if divisorSum == n {
		return Perfect
	} else if divisorSum > n {
		return Abundant
	}
	return Deficient
}

// sumOfProperDivisors calculates the sum of all divisors of n except n itself
func sumOfProperDivisors(n int) int {
	if n <= 1 {
		return 0
	}

	sum := 1 // 1 is always a proper divisor for n > 1

	// Only need to check up to sqrt(n) for efficiency
	// For each divisor i found, n/i is also a divisor
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			// Add the complementary divisor if it's different from i
			if i != n/i {
				sum += n / i
			}
		}
	}

	return sum
}
