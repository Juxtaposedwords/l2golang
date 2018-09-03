package remove

import (
	"testing"
)

var (
	RemoveElement = removeElement
)

func TestRemoveElement(t *testing.T) {
	type h struct {
		nums []int
		val  int
	}
	tt := []struct {
		have *h
		want int
	}{
		{&h{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
		{&h{[]int{3, 2, 2, 3}, 3}, 2},
	}

	for _, v := range tt {
		got := removeElement(v.have.nums, v.have.val)
		if got != v.want {
			t.Errorf("removeElement: %#v got: %d want: %d\n", v.have, got, v.want)
		}
	}
}
