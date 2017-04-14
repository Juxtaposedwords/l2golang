package tree

import (
	"testing"
)

func TestInsert(t *testing.T) {

	root := &Node{key: 5}
	root.Insert(4)
	root.Insert(3)
	root.Insert(2)
	root.Insert(1)
	root.Insert(4)
	root.Insert(6)
	root.Insert(7)
	root.Insert(3)
	root.Traverse()
	//fmt.Printf("%v\n", root)
/*	x := *root.left
	if x.key != 3 {
		t.Errorf("Insertfailed: got %d, want 3 ", root.left.key)
	} */
}