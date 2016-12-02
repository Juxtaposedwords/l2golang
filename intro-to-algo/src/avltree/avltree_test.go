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
}