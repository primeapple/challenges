package shuffleLine

import "slices"

func ShuffleLine(input []string, n int) []string {
	start := []string{}
	end := []string{}

	for idx, name := range input {
		if (idx+1)%n == 0 {
			end = append(end, name)
		} else {
			start = append(start, name)
		}
	}

	return slices.Concat(start, end)
}
