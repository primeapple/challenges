package sortHtmlColors

import (
	"slices"
)

func SortHtmlColors(colors []string) []string {
	clonedColors := slices.Clone(colors)

	for i := 0; i < len(clonedColors); i++ {
		for j := 0; j < len(clonedColors); j++ {
			if CompareString(clonedColors[i], clonedColors[j]) < 0 {
				clonedColors[i], clonedColors[j] = clonedColors[j], clonedColors[i]
			}
		}
	}

	return clonedColors
}

// Compares two strings, returning positive value if the first one is "bigger" the second one, negative value otherwise
func CompareString(s1, s2 string) int {
	rune1 := []rune(s1)
	rune2 := []rune(s2)

	for i := 0; i < max(len(rune1), len(rune2)); i++ {
		if i == len(rune1) {
			return -1
		}
		if i == len(rune2) {
			return 1
		}

		if int(rune1[i]) != int(rune2[i]) {
			return int(rune1[i]) - int(rune2[i])
		}
	}

	return 0
}
