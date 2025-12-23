package alternatingArray

func AlternatingArrayAI(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	even := arr[0]
	odd := arr[1]
	for i, v := range arr {
		if i%2 == 0 {
			if v != even {
				return false
			}
		} else {
			if v != odd {
				return false
			}
		}
	}
	return true
}
