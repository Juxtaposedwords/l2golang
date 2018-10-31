package rotate

import (
	"testing"
)

var RotateRight = rotateRight

func TestRotateRight(t *testing.T) {

type h struct {
    node *ListNode
    rotate int
}
	tt := []struct {
		have *h
		want *ListNode
	}{
		{
		  &h{&ListNode{0, &{ListNode{1, &ListNode{2,&ListNode{3,nil}}}}},1},
          &ListNode{3, &{ListNode{0, &ListNode{1,&ListNode{2,nil}}}}}
		},
		{[2]string{"1010", "1011"}, "10101"},
		{[2]string{"1111", "1111"}, "11110"},
	}

	for _, v := range tt {
		got := addBinary(v.have[0], v.have[1])
		if got != v.want {
			t.Errorf("\nlongestPalindrome failed with\n  test: %#v\n  got: %#v\n", v, got)
		}
	}
}
