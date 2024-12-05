package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Comparator struct {
	Smaller []int
	Bigger  []int
}

func (c *Comparator) AddSmaller(x int) {
	c.Smaller = append(c.Smaller, x)
}
func (c *Comparator) AddBigger(x int) {
	c.Bigger = append(c.Bigger, x)
}

type Ruleset map[int]*Comparator

func (r Ruleset) AddComparison(left, right int) {
	leftComparator, found := r[left]
	if found {
		leftComparator.AddBigger(right)
	} else {
		r[left] = &Comparator{[]int{}, []int{right}}
	}

	rightComparator, found := r[right]
	if found {
		rightComparator.AddSmaller(left)
	} else {
		r[right] = &Comparator{[]int{left}, []int{}}
	}
}

func (r Ruleset) IsCorrectlyOrdered(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		comparator, found := r[numbers[i]]
		if !found {
			continue
		}

		for j := i + 1; j < len(numbers); j++ {
			if slices.Contains(comparator.Smaller, numbers[j]) {
				return false
			}
		}
	}
	return true
}

func (r Ruleset) Compare(num1, num2 int) int {
	comparator, found := r[num1]
	if !found {
		return 0
	}

	if slices.Contains(comparator.Bigger, num2) {
		return -1
	} else if slices.Contains(comparator.Smaller, num2) {
		return 1
	}

	return 0
}

func CreateRuleset(lines []string) Ruleset {
	ruleset := Ruleset{}
	for _, line := range lines {
		rule := strings.Split(line, "|")
		if len(rule) != 2 {
			panic(fmt.Sprintf("unexpected rule found %v", rule))
		}
		left, err := strconv.Atoi(rule[0])
		if err != nil {
			panic(err.Error())
		}
		right, err := strconv.Atoi(rule[1])
		if err != nil {
			panic(err.Error())
		}
		ruleset.AddComparison(left, right)
	}

	return ruleset
}

func SumMiddlePageOfOrdered(lines []string, r Ruleset) int {
	sum := 0
	for _, line := range lines {
		stringNums := strings.Split(line, ",")
		nums := make([]int, 0, len(stringNums))
		for _, stringNum := range stringNums {
			num, err := strconv.Atoi(stringNum)
			if err != nil {
				panic(fmt.Sprintf("Got unexpected non int %q", stringNum))
			}
			nums = append(nums, num)
		}

		if r.IsCorrectlyOrdered(nums) {
			sum += nums[len(nums)/2]
		}
	}

	return sum
}

func SumMiddlePageOfUnordered(lines []string, r Ruleset) int {
	sum := 0
	for _, line := range lines {
		stringNums := strings.Split(line, ",")
		nums := make([]int, 0, len(stringNums))
		for _, stringNum := range stringNums {
			num, err := strconv.Atoi(stringNum)
			if err != nil {
				panic(fmt.Sprintf("Got unexpected non int %q", stringNum))
			}
			nums = append(nums, num)
		}

		if !r.IsCorrectlyOrdered(nums) {
            slices.SortFunc(nums, r.Compare)
			sum += nums[len(nums)/2]
		}
	}

	return sum
}

func main() {
	body, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(body), "\n\n")
	rules := strings.Split(parts[0], "\n")
	pages := strings.Split(parts[1], "\n")
	pages = pages[:len(pages)-1]

	ruleset := CreateRuleset(rules)

	solution1 := SumMiddlePageOfOrdered(pages, ruleset)
	solution2 := SumMiddlePageOfUnordered(pages, ruleset)
	fmt.Println("Solution puzzle 1:", solution1)
	fmt.Println("Solution puzzle 2:", solution2)
}
