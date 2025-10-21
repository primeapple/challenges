package splitbywidths

func SplitByWidths(str string, widths []int) []string {
	if len(str) == 0 || len(widths) == 0 {
		return []string{}
	}

	strLen := len(str)

	substrings := []string{}
	start := 0

	slice := func(width int) int {
		end := min(strLen, start+width)
		substrings = append(substrings, str[start:end])
		return end
	}

	for _, width := range widths {
		start = slice(width)
	}

	lastWidth := widths[len(widths)-1]
	for start != strLen {
		start = slice(lastWidth)
	}

	return substrings
}
