package duplicates

/*
This is kind of a wacky problem. So here's the prompt verbatim. Typically I try
to distil the problem.

Given an ordered slice of integers:
    * order the passed slice so the first items are all unique is ascending order
    * return the length of items

To accomplish this we create a struct to store the value and last index.
*/
type last struct {
	Val   int
	Index int
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l := &last{Val: nums[0], Index: 0}
	for i := 1; i < len(nums); i++ {
		if l.Val != nums[i] {
			l.Val = nums[i]
			nums[l.Index+1], nums[i] = nums[i], nums[l.Index+1]
			l.Index++
		}
	}
	return l.Index + 1
}
