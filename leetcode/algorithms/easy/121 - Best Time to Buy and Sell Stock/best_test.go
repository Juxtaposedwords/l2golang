package best

import (
	"testing"
)

var (
	MaxProfit = maxProfit
)

func MaxProfit_test(t *testing.T) {
	tt := []struct {
		have []int
		want int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, v := range tt {
		got := maxProfit(v.have)
		if got != v.want {
			t.Errorf("%#v %#v %#v", v.have, v.want, got)
		}
	}
}
