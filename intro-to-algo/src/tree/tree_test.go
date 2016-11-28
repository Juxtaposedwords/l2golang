package tree

import (
	"testing"
	"fmt"
)

func TestInsert(t *testing.T) {
	root := &Node{key: 5}
	root.Insert(3)
	fmt.Printf("%z\n", root)
/*	x := *root.left
	if x.key != 3 {
		t.Errorf("Insertfailed: got %d, want 3 ", root.left.key)
	} */
}