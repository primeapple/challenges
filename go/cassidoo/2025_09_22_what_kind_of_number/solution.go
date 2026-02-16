package whatKindOfNumber

import (
	"cassidoo/utils"
)

func WhatKindOfNumber(n int) KindOfNumber {
	divisors := getProperDivisors(n)

	properDivisorsSum := utils.Reduce(divisors, 0, func(next, acc int) int {
		return next + acc
	})

	if properDivisorsSum == n {
		return Perfect
	} else if properDivisorsSum > n {
		return Abundant
	}
	return Deficient
}

func getProperDivisors(n int) []int {
	divisors := []int{1}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
