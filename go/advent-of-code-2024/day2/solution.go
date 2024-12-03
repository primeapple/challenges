package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func boolPointer(b bool) *bool {
	return &b
}

func isSafe(report []int) bool {
	var lastLevel *int
	var isAscending *bool

	for _, level := range report {
		if lastLevel != nil {
			if abs(level-*lastLevel) > 3 {
				return false
			}

			if isAscending != nil {
				if *isAscending && level <= *lastLevel {
					return false
				}
				if !*isAscending && level >= *lastLevel {
					return false
				}
			} else {
				if level > *lastLevel {
					isAscending = boolPointer(true)
				} else if level < *lastLevel {
					isAscending = boolPointer(false)
				} else {
					return false
				}
			}
		}
		lastLevel = &level
	}

	return true
}

func GenerateSlices(fullArray []int) [][]int {
	var slices [][]int
    for i := range fullArray {
        newSlice := append([]int{}, fullArray[:i]...)
        newSlice = append(newSlice, fullArray[i+1:]...)
		slices = append(slices, newSlice)
	}
	return slices
}

func ReportSafetyChecking(reports [][]int, tolerateSingleBadLevel bool) int {
	numSafeReports := 0

	for _, report := range reports {
		isSafeFirstTry := isSafe(report)
		if isSafeFirstTry {
			numSafeReports++
		} else if tolerateSingleBadLevel {
			for _, slice := range GenerateSlices(report) {
				isSafeInSlice := isSafe(slice)
				if isSafeInSlice {
					numSafeReports++
					break
				}
			}
		}
	}

	return numSafeReports
}

func lineToInts(line string) ([]int, error) {
	stringNumbers := strings.Split(line, " ")
	var numbers = make([]int, 0, len(stringNumbers))

	for _, number := range stringNumbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func main() {
	body, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(body), "\n")

	var reports [][]int

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}

		report, err := lineToInts(lines[i])
		if err != nil {
			panic(err)
		}

		reports = append(reports, report)
	}

	solution1 := ReportSafetyChecking(reports, false)
	fmt.Println("Solution puzzle 1:", solution1)

	solution2 := ReportSafetyChecking(reports, true)
	fmt.Println("Solution puzzle 2:", solution2)
}
