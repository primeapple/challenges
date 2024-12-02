package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func DistanceSumBetweenLists(list1, list2 []int) (int, error) {
	if len(list1) != len(list2) {
		return 0, errors.New("lists should be identical")
	}

	sortedList1 := append([]int{}, list1...)
	sortedList2 := append([]int{}, list2...)
	slices.Sort(sortedList1)
	slices.Sort(sortedList2)

	distanceSum := 0
	for i := 0; i < len(sortedList1); i++ {
		distanceSum += abs(sortedList1[i] - sortedList2[i])
	}

	return distanceSum, nil
}

func SimilarityScore(list1, list2 []int) int {
    occurrencesInList2ByNumber := make(map[int]int)
	for i := 0; i < len(list2); i++ {
        occurence, exists := occurrencesInList2ByNumber[list2[i]]
        if exists {
            occurrencesInList2ByNumber[list2[i]] = occurence + 1
        } else {
            occurrencesInList2ByNumber[list2[i]] = 1
        }
	}

	similarityScore := 0
	for _, number := range list1 {
		occurrence, ok := occurrencesInList2ByNumber[number]
		if !ok {
			continue
		}
		similarityScore += number * occurrence
	}

    return similarityScore
}

func lineToInts(line string) (int, int, error) {
	numbers := strings.Split(line, "   ")
	if len(numbers) != 2 {
		return 0, 0, fmt.Errorf("Not exactly 2 numbers found in line %q", line)
	}

	num1, err := strconv.Atoi(numbers[0])
	if err != nil {
		return 0, 0, err
	}

	num2, err := strconv.Atoi(numbers[1])
	if err != nil {
		return 0, 0, err
	}

	return num1, num2, nil
}

func main() {
	body, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(body), "\n")

	var list1 []int
	var list2 []int

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}

		num1, num2, err := lineToInts(lines[i])
		if err != nil {
			panic(err)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	solution1, err := DistanceSumBetweenLists(list1, list2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Solution puzzle 1:", solution1)

	solution2 := SimilarityScore(list1, list2)
	fmt.Println("Solution puzzle 2:", solution2)
}
