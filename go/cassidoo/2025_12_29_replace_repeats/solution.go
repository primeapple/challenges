package replaceRepeats

import (
	"fmt"
	"strconv"
	"strings"
)

func ReplaceRepeats(digits string, n int) (string, error) {
	var builder strings.Builder

	consecutiveCount := 0
	for _, digit := range digits {
		intDigit, err := strconv.Atoi(string(digit))
		if err != nil {
			return "", fmt.Errorf("couldn't convert to number: %d (%v): %w", digit, err, ErrInvalidDigit)
		}

		if intDigit == n {
			consecutiveCount++
		} else {
			if consecutiveCount > 0 {
				builder.WriteString(strconv.Itoa(consecutiveCount))
				consecutiveCount = 0
			}
			builder.WriteRune(digit)
		}
	}

	if consecutiveCount > 0 {
		builder.WriteString(strconv.Itoa(consecutiveCount))
	}

	return builder.String(), nil
}
