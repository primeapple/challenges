package longestCoprimeSubsequence

import (
	"fmt"
)

func LongestCoprimeSubsequence(numbers []int) int {
	subsequences := Subsequences(numbers)

	totalLongestSubsequence := 1
	for _, p := range subsequences {
		currentSubSequenceLength := longestCoprimeSubsequencePerArray(p)
		totalLongestSubsequence = max(totalLongestSubsequence, currentSubSequenceLength)
		if currentSubSequenceLength == 6 {
			fmt.Printf("%v", p)
		}
	}
	return totalLongestSubsequence
}

func Subsequences(numbers []int) [][]int {
	if len(numbers) == 1 {
		return [][]int{numbers}
	}
	subsequences := [][]int{}
	nextSubsequences := Subsequences(numbers[1:])
	for _, nextSubsequence := range nextSubsequences {
		subsequences = append(subsequences, append([]int{numbers[0]}, nextSubsequence...))
		subsequences = append(subsequences, nextSubsequence)
	}
	return subsequences
}

func longestCoprimeSubsequencePerArray(numbers []int) int {
	totalLongestSubsequence := 1
	currentLongestSubsequence := 1
	for i := 1; i < len(numbers); i++ {
		if isCoprime(numbers[i-1], numbers[i]) {
			currentLongestSubsequence++
		} else {
			totalLongestSubsequence = max(totalLongestSubsequence, currentLongestSubsequence)
			currentLongestSubsequence = 1
		}
	}
	totalLongestSubsequence = max(totalLongestSubsequence, currentLongestSubsequence)

	return totalLongestSubsequence
}

func isCoprime(a, b int) bool {
	if a%2 == 0 && b%2 == 0 {
		return false
	}
	for i := 3; i <= min(a, b); i += 2 {
		if a%i == 0 && b%i == 0 {
			return false
		}
	}
	return true
}
