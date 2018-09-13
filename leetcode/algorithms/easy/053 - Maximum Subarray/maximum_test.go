package maximum

import (
	"testing"
)

var MaxSubArray = maxSubArray

func TestMaxSubArray(t *testing.T) {

	tt := []struct {
		have []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{0, 0, -21, 0}, 0},
		{[]int{-32, -21, -100, -2099, -212}, -21},
	}

	for _, v := range tt {
		got := maxSubArray(v.have)
		if got != v.want {
			t.Errorf("\nlongestPalindrome failed with\n  test: %#v\n  got: %#v\n", v, got)
		}
	}
}
