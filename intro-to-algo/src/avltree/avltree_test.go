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
  		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 2, 3, 3, 4, 4, 5, 7}, "Traverse Works"},
  	//	{[]int{5, 4, 3, 3, 2, 2, 1}, []int{1, 2, 2, 3, 3, 4, 5}, "reversed"},
	}
	for _, test := range(tcs){
		tree := &Tree{} 
		for _, i := range(test.have){
			tree.Insert(i)
		}
		//tree.Rebalance()
		tree.Traverse()
		tree.left_rotate()
		fmt.Printf("DID A ROTATE\n\n")
		fmt.Printf("The root is: %z\n", tree.root)
		tree.Traverse()

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