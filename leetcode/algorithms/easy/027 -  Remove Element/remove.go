package remove

/*

Prompt:
Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Explanation:
This is a classic 'swap' example.

Go through each character. If it isn't the offending character, 'val', swap it
with the last known good point


*/
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
