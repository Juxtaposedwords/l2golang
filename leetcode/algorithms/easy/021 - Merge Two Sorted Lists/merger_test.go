package merger

import (
	"fmt"
	"reflect"
	"testing"
)

var MergeTwoLists = mergeTwoLists

func TestMergeTwoLists(t *testing.T) {

	tt := []struct {
		have []*ListNode
		want *ListNode
	}{
		{
			[]*ListNode{
				&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
				&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
			},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, &ListNode{5, &ListNode{6, &ListNode{6, nil}}}}}}}}}}}},
		},
		{
			[]*ListNode{
				&ListNode{0, nil},
				&ListNode{0, nil},
			},
			&ListNode{0, &ListNode{0, nil}},
		},
	}
	for _, v := range tt {
		got := llToSlice(mergeTwoLists(v.have[0], v.have[1]))
		if !reflect.DeepEqual(got, llToSlice(v.want)) {
			t.Errorf("mergeTwoLists\n v:%#v\n want: %#v\n got:  %#v\n ", v, v.want, got)
		}
	}
}

func sliceToLL(input []int) *ListNode {
	fmt.Printf("SlicetoLL: %#v\n", input)
	f := &ListNode{Val: input[0]}
	o := f
	for i := 1; i < len(input); i++ {
		f.Next = &ListNode{Val: input[i]}
		f = f.Next
	}
	return o

}

func llToSlice(input *ListNode) []int {
	o := []int{}
	for {
		if input == nil {
			return o
		}
		o = append(o, input.Val)
		input = input.Next
	}
}
