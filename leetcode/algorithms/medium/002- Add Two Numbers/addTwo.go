package addTwo

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// This exercise could be done in a more clean manner with the container/list
//    package, but i wanted to practice with handling lists given their
//    relation to trees and tries.

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return UInt64ToList(ListToUInt64(l1) + ListToUInt64(l2))
}
func ListToUInt64(ll *ListNode) uint64 {
	node := ll
	d := []string{}
	for {
		d = append(d, strconv.Itoa(node.Val))
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	// Now let's get everything in the right order
	d = reverseArray(d)
	o, _ := strconv.ParseUint(strings.Join(d, ""), 10, 64)
	return o
}
func UInt64ToList(input uint64) *ListNode {
	// Split  the UInt64eger UInt64o a slice for each digit
	d := strings.Split(fmt.Sprintf("%d", input), "")

	// set the UInt64ial node
	x, _ := strconv.Atoi(d[len(d)-1])
	node := &ListNode{Val: x}
	// store a poUInt64er to the first node for a return value
	f := node
	// go through the slice of UInt64 digits backwards

	for i := len(d) - 2; i >= 0; i-- {
		// convert the string digit UInt64o an UInt64 and disregard the error
		x, _ := strconv.Atoi(d[i])
		// creat the next node
		node.Next = &ListNode{}
		// swap out of the next node
		node = node.Next
		// Set the node's value to an UInt64eger
		node.Val = x
	}
	return f
}
func reverseArray(input []string) []string {
	for i := len(input)/2 - 1; i >= 0; i-- {
		opp := len(input) - 1 - i
		input[i], input[opp] = input[opp], input[i]
	}
	return input
}
