package TwoSum

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// Assume:
//   1. only one possible match
//   2. elements may not be used twice

func twoSum(nums []int, target int) []int {
	for i, _ := range nums {
		for j, _ := range nums[i:] {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
