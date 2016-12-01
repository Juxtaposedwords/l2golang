package avltree

import (
	"testing"
)

func TestInsert(t *testing.T) {

	root := &Node{key: 1}
	root.Insert(2)
	root.Insert(3)
	root,_ = left_rotate(root)
	root.Traverse()
	//fmt.Printf("%v\n", root)
/*	x := *root.left
	if x.key != 3 {
		t.Errorf("Insertfailed: got %d, want 3 ", root.left.key)
	} */
}