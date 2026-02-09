package moveNums

// MoveNumsAI moves all occurrences of n to the end of the array
// while maintaining the relative order of non-n elements.
// This is done in-place without creating a copy of the array.
func MoveNumsAI(nums []int, n int) []int {
	// Use two pointers: writePos tracks where to write non-n elements
	writePos := 0

	// First pass: move all non-n elements to the front
	for readPos := 0; readPos < len(nums); readPos++ {
		if nums[readPos] != n {
			nums[writePos] = nums[readPos]
			writePos++
		}
	}

	// Second pass: fill remaining positions with n
	for i := writePos; i < len(nums); i++ {
		nums[i] = n
	}

	return nums
}
