package moveNums

func MoveNums(nums []int, n int) []int {
	lastIndexToCheck := len(nums) - 1
	for i := 0; i <= lastIndexToCheck; i++ {
		value := nums[i]
		if value == n {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, value)
			lastIndexToCheck--
		}
	}
	return nums
}
