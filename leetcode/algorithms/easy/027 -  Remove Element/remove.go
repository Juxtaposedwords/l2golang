package remove

func removeElement(nums []int, val int) int {
	next := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[next], nums[i] = nums[i], nums[next]
			next++
		}
	}
	nums = nums[:next]
	return next
}
