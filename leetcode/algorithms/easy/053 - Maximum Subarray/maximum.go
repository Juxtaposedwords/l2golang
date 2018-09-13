package maximum

/*

At a first glance we can probably solve this in N time,  talking later pointed
out that since this is an order list it could possibly be done via divide
and conquer meaning there may be a nlogn solution.


Notable characteristics of the problem:
* There is no state or specifics required. We don't have to know which nodes,
just that there is a largest number.

N time solution:
1. Go through each number and keep a running max
2.


nlogn time soultion:

What do we divide on?

*/

func maxSubArray(nums []int) int {
	return nMaxSubArray(nums)
}

func nMaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, localMax := nums[0], nums[0]
	for _, e := range nums[1:len(nums)] {
		if localMax+e < e {
			localMax = e
		} else {
			localMax += e

		}
		if localMax > max {
			max = localMax
		}
	}
	return max
}
