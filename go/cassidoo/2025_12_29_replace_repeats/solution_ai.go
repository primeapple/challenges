package replaceRepeats

import (
	"fmt"
	"strconv"
)

func ReplaceRepeatsAI(digits string, n int) (string, error) {
	result := ""
	i := 0
	for i < len(digits) {
		if digits[i] < '0' || digits[i] > '9' {
			return "", fmt.Errorf("invalid digit %c: %w", digits[i], ErrInvalidDigit)
		}
		digit := int(digits[i] - '0')
		if digit == n {
			count := 0
			for i < len(digits) && int(digits[i]-'0') == n {
				count++
				i++
			}
			result += strconv.Itoa(count)
		} else {
			result += string(digits[i])
			i++
		}
	}
	return result, nil
}
