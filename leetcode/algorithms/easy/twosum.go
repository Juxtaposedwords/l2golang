package TwoSum

import "fmt"

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// Assume:
//   1. only one possible match
//   2. elements may not be used twice

// My solution:
//    Steps through the arraywith a sub array, meaning we have N^2.
func twoSum(nums []int, target int) []int {
	fmt.Println("new")

	for i, _ := range nums {
		for j, _ := range nums[i+1:] {
			fmt.Printf("i: %d, j: %d\n", i, j)
			if nums[i]+nums[j+i+1] == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return []int{}
}
