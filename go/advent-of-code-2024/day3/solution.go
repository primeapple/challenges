package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Pair struct {
	first  int
	second int
}

func (p Pair) Result() int {
	return p.first * p.second
}

func ParseNumbers(text string, useDoDont bool) []Pair {
	var parsed []Pair
	re1 := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	re2 := regexp.MustCompile(`(?:do\(\))|(?:don't\(\))|(?:mul\((\d{1,3}),(\d{1,3})\))`)

	var matches [][]string
	if useDoDont {
		matches = re2.FindAllStringSubmatch(text, -1)
	} else {
		matches = re1.FindAllStringSubmatch(text, -1)
	}

	isEnabled := true
	for _, match := range matches {
		if useDoDont && match[0] == "do()" {
			isEnabled = true
            continue
		} else if match[0] == "don't()" {
			isEnabled = false
            continue
		}

		if useDoDont && !isEnabled {
			continue
		}

		firstMatchNumber := match[1]
		secondMatchNumber := match[2]

		num1, err := strconv.Atoi(firstMatchNumber)
		if err != nil {
			panic("Cant convert strin to int")
		}
		num2, err := strconv.Atoi(secondMatchNumber)
		if err != nil {
			panic("Cant convert strin to int")
		}
		parsed = append(parsed, Pair{num1, num2})
	}

	return parsed
}

func main() {
	body, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	numbers1 := ParseNumbers(string(body), false)
	numbers2 := ParseNumbers(string(body), true)
	solution1 := 0
	solution2 := 0
	for _, num := range numbers1 {
		solution1 += num.Result()
	}
	for _, num := range numbers2 {
		solution2 += num.Result()
	}
	fmt.Println("Solution puzzle 1:", solution1)
	fmt.Println("Solution puzzle 2:", solution2)
}
