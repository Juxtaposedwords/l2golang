package addTwo

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tt := []struct {
		have [2]*ListNode
		want *ListNode
	}{
		{
			[2]*ListNode{
				&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
				&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
	}
	for _, v := range tt {
		got := AddTwoNumbers(v.have[0], v.have[1])
		got_slice, want_slice := nodeToSlice(got), nodeToSlice(v.want)

		if !reflect.DeepEqual(got_slice, want_slice) {
			t.Errorf("Error with AddTwoNumbers: %#V, got %v want : %v", v, got_slice, want_slice)
		}

	}
}

func TestListToUInt64(t *testing.T) {

	tt := []struct {
		have *ListNode
		want int
	}{
		{&ListNode{9, &ListNode{8, &ListNode{7, nil}}}, 789},
	}
	for _, v := range tt {
		got := strings.Split(strconv.Itoa(ListToUInt64(v.have)), "")
		want := nodeToSlice(v.have)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error with ListToUInt64: %#V, got %v want : %v", v, got, want)
		}

	}
}
func TestUInt64ToList(t *testing.T) {
	tt := []struct {
		have int
		want *ListNode
	}{
		{789, &ListNode{9, &ListNode{8, &ListNode{7, nil}}}},
		{777, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}},
		{0, &ListNode{0, nil}},
	}
	for _, v := range tt {
		got := nodeToSlice(UInt64ToList(v.have))
		want := strings.Split(strconv.Itoa(v.have), "")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error with UInt64ToList: %#V, got %v want : %v", v, got, want)
		}
	}
}

func nodeToSlice(input *ListNode) []string {
	got_UInt64s := []string{}
	node := input
	for {
		got_UInt64s = append(got_UInt64s, strconv.Itoa(node.Val))
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return reverseArray(got_UInt64s)
}
