package majorityElement

import "fmt"

func MajorityElement(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("numbers are empty")
	}
	count := 0
	element := numbers[0]

	for _, number := range numbers {
		if count == 0 {
			count++
			element = number
		} else if element == number {
			count++
		} else {
			count--
		}
	}

	totalCount := 0
	for _, number := range numbers {
		if number == element {
			totalCount++
		}
	}

	if totalCount <= len(numbers)/2 {
		return 0, ErrNoMajorityFound
	}

	return element, nil
}
