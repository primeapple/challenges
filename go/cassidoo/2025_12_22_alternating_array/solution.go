package alternatingArray

func AlternatingArray(array []int) bool {
	if len(array) <= 2 {
		return true
	}

	even, odd := array[0], array[1]

	for i := 2; i < len(array); i = i + 2 {
		if even != array[i] {
			return false
		}
	}

	for i := 3; i < len(array); i = i + 2 {
		if odd != array[i] {
			return false
		}
	}

	return true
}
