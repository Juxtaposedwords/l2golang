package avltree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := &Tree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.root.Traverse()
}