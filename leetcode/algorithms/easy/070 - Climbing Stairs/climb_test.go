package climb

import (
	"testing"
)

var ClimbStairs = climbStairs

func TestClimbStairs(t *testing.T) {

	tt := []struct {
		have int
		want int
	}{
		{1, 1},  //
		{2, 2},  // 1
		{3, 3},  // 1
		{4, 5},  // 2
		{5, 8},  // 3
		{6, 13}, // 5
		{7, 21}, // 8
		{8, 34}, //11
		{9, 55}, //21
	}

	for _, v := range tt {
		got := climbStairs(v.have)
		if got != v.want {
			t.Errorf("\nlongestPalindrome failed with\n  test: %#v\n  got: %#v\n", v, got)
		}
	}
}
