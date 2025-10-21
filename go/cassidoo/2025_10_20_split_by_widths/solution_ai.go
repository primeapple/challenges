package splitbywidths

func SplitByWidthsAI(str string, widths []int) []string {
	if len(str) == 0 || len(widths) == 0 {
		return []string{}
	}

	result := []string{}
	pos := 0
	widthIdx := 0

	for pos < len(str) {
		width := widths[widthIdx]

		end := pos + width
		if end > len(str) {
			end = len(str)
		}

		result = append(result, str[pos:end])
		pos = end

		if widthIdx < len(widths)-1 {
			widthIdx++
		}
	}

	return result
}
