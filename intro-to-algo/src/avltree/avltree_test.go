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
  		{[]int{1,2,3,4 }, []int{1, 2, 2, 3, 3, 4, 4, 5, 7}, "Traverse Works"},
  	//	{[]int{5, 4, 3, 3, 2, 2, 1}, []int{1, 2, 2, 3, 3, 4, 5}, "reversed"},
	}
	for _, test := range(tcs){
		tree := &Tree{} 		
		f := func(n *Node) {
			var s string
			for i:=0; i <n.depth; i ++{
				s +="="
			}
			//fmt.Printf("%skey: %d\n  depth:%d\n  height:%d\n  left: %z\n  right: %z\n",s, n.key, n.depth,n.height, n.left, n.right)
			fmt.Printf("%s>(%d)\n",s,n.key)
		}
		for _, i := range(test.have){
			tree.Insert(i)
			tree.Traverse(f)
			fmt.Printf("----\n")

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

func TestLeftRotate(t *testing.T){
	root := &Node{key: 2}
	root.left = &Node{key: 1}
	root.right = &Node{key: 3}
	a, b :=  root.left, root
	root.left_rotate()
	root = a
	if root.left != nil && root.right != b {
		t.Errorf("Your left rotate is broken.")	
	}

}
func TestRightRotate(t *testing.T){
	root := &Node{key: 2}
	root.left = &Node{key: 1}
	root.right = &Node{key: 3}
	a, b :=  root.right, root
	root.right_rotate()
	root = a
	if root.left != nil && root.left != b {
		t.Errorf("Your left rotate is broken.")	
	}
}