package perrinCombinations

import (
	"cassidoo/utils"
	"slices"
)

func PerrinCombinations(n, k int) [][]int {
	uniquePerrinNumbers := utils.Unique(perrinNumbers(n))
	slices.Sort(uniquePerrinNumbers)

	return perrinNumbersSum(0, k, uniquePerrinNumbers, []int{})
}

func perrinNumbersSum(current, wanted int, remainingNumbers, takenNumbers []int) [][]int {
	if current == wanted {
		return [][]int{takenNumbers}
	}
	if len(remainingNumbers) == 0 {
		return [][]int{}
	}
	if current > wanted {
		return [][]int{}
	}
	newRemainingNumbers := slices.Clone(remainingNumbers)
	number, _ := utils.Pop(&newRemainingNumbers)
	return append(perrinNumbersSum(current+number, wanted, newRemainingNumbers, append(takenNumbers, number)), perrinNumbersSum(current, wanted, newRemainingNumbers, takenNumbers)...)
}

func perrinNumbers(n int) []int {
	perrinNumbers := make([]int, n)
	for i := range n {
		switch i {
		case 0:
			perrinNumbers[0] = 3
		case 1:
			perrinNumbers[1] = 0
		case 2:
			perrinNumbers[2] = 2
		default:
			perrinNumbers[i] = perrinNumbers[i-2] + perrinNumbers[i-3]
		}
	}

	return perrinNumbers
}
