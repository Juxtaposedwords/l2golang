package avltree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	 tcs := []struct{
		have []int
		want []int
		desc string
	}{
  		{[]int{1, 2, 3,4,4,2, 5, 6 }, []int{1, 2, 2, 3, 3, 4, 4, 5, 7}, "Traverse Works"},
  	//	{[]int{5, 4, 3, 3, 2, 2, 1}, []int{1, 2, 2, 3, 3, 4, 5}, "reversed"},
	}
	for _, test := range(tcs){
		tree := &Tree{} 
		for _, i := range(test.have){
			tree.Insert(i)
		}
		f := func(n *Node) {
			var s string
			for i:=0; i <n.depth; i ++{
				s +=" "
			}
			fmt.Printf("%s%z\n", s, n)
			//fmt.Printf("%d\n  depth:%d\n  height:%d\n  left: %z\n  right: %z\n", n.key, n.depth,n.height, n.left, n.right)
		}
		tree.Rebalance()
		tree.Traverse(f)

	}


}
func TestNodeInsert(t *testing.T){
	n := &Node{ key: 1}
	if n.right != nil {
		t.Errorf("Node Class: Left key should be nil on instantiation")
	}
	n.Insert(2)
	if n.left != nil {
		t.Errorf("Node Class: Left key should be nil on instantiation")
	}
	n.Insert(1)
	if n.right.key != 2 {
		t.Errorf("Node Class: Failed to insert correctly to the right")
	}
	if n.left.key != 1 {
		t.Errorf("Node Class: Faield to insert correctly to the left")	
	}
	n.Insert(1)
	if n.left.left.key != 1 {
		t.Errorf("Node Class: Failed to insert correctly to the left's child left value")	
	}
}