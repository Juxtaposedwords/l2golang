package merger

import (
//	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	f := &ListNode{}
	o := f
	for {
		switch {
		case l1 != nil && l2 != nil && l1.Val >= l2.Val:
			o.Val = l2.Val
			l2 = l2.Next
		case l1 != nil && l2 != nil && l1.Val < l2.Val:
			o.Val = l1.Val
			l1 = l1.Next
		case l1 == nil:
			o.Val = l2.Val
			l2 = l2.Next
		case l2 == nil:
			o.Val = l1.Val
			l1 = l1.Next
		}
		if l1 == nil && l2 == nil {
			break
		}
		o.Next = &ListNode{}
		o = o.Next

	}
	return f
}
