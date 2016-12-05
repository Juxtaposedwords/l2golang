package avltree
import (
//	"fmt"
)
type Node struct{
	key int
	height int
	left *Node
	right *Node
}
type Tree struct{
	root *Node
}

func(t *Tree) Insert(k int){
	if t.root == nil{
		t.root = &Node{key: k, height: 0}
	} else {
		t.root.Insert(k)
		t.Rebalance()
	}

}
func(n *Node) Insert(k int){
	switch {
	case k <= n.key && n.left == nil : 
		n.left = &Node{key: k, height: 1}
	case k > n.key  && n.right == nil :
		n.right = &Node{key: k, height: 1}
	case k <= n.key  && n.left != nil : 
		n.left.Insert(k)
	case k > n.key  && n.right != nil :
		n.right.Insert(k)
	}
	n.SetHeight()
}
func(n *Node) SetHeight(){
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}

	if left <= right {
		n.height = right + 1
	} else if left > right {
		n.height = left + 1
	}	
}
func(n *Node) Traverse(f func(*Node))  {
	if n.left != nil {
		n.left.Traverse(f)
	}
	f(n)
	if n.right != nil {
		n.right.Traverse(f)
	}
}
func(t *Tree) Traverse(f func(*Node)) {
	t.root.Traverse(f)
}
func(t *Tree) Rebalance() {
	left, right := 0,0
	if t.root.left != nil {
		left = t.root.left.height
	} 
	if t.root.right != nil{
		right = t.root.right.height 
	}
	diff := left - right
	switch {
	case diff < -1:
		t.left_rotate()
		t.Rebalance()
	case diff > 1: 
		t.right_rotate()
		t.Rebalance()
	}	

}

// Left rotates move the previously right child node to being the 
// root node
//  1. Store the new and old root as n and o
//  2. Store the left child (l) of the new root, as if it is not nil it will need to be preserved

func(t *Tree) left_rotate(){
	n := t.root.right
	o := t.root
	l := n.left
	n.left = o
	o.right = l
	t.root = n
}
// Right Rotates move the previously left child node to being the
// root node
//  1. Store the new and old root as n and o
//  2. Store the right child(r) of the new root, as if it is not nil it will need to be preserved
func(t *Tree) right_rotate(){ 
	n := t.root.left
	o := t.root
	r := n.right
	n.right = o
	o.left = r
	t.root = n
}