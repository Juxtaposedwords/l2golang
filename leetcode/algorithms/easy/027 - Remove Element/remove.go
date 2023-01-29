package remove

func Remover(nums []int, val int) int {
	j := 0
	for i, number := range nums {
		// While an early exit, this parses more as human would solve the problem.
		if number == val {
			continue
		}
		if j < i {
			nums[j] = number
		}
		j++
	}
	return j
}
