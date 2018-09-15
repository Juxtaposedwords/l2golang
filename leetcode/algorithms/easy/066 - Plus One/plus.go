package plus

func plusOne(digits []int) []int {
	carry := false
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			carry = true
			digits[i] = 0
		} else {
			digits[i] += 1
			carry = false
		}

		if !carry {
			return digits
		}
	}
	return append([]int{1}, digits...)
}
