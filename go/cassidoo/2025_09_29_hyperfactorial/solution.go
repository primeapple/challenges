package hyperfactorial

import "math"

func Hyperfactorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * int(math.Pow(float64(i), float64(i)))
	}
	return result
}
