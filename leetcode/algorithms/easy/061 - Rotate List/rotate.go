package binary

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	stack := []*Listnode{head}
	var last, iter *ListNode
	iter = head
	for iter != nil {
		if len(stack) >= k+1 {
			stack = stack[:0]
		}
		stack = append(stack, iter)
		last = iter
		iter = iter.Next
	}
	last.Next = head
	stack[0].Next = nil

}

func reverseNodeSlice(input []ListNode) []ListNode {
	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-i] = input[len(input)-i], input[i]
	}
}
