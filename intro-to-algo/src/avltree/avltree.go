package avltree
import (
//	"fmt"
)
type Node struct{
	key int
	height int
	depth int
	left *Node
	right *Node
}
type Tree struct{
	root *Node
}

func(t *Tree) Insert(k int){
	if t.root == nil{
		t.root = &Node{key: k, height: 0, depth: 1}
	} else {
		t.root.Insert(k)
		//t.Rebalance()
	}

}
func(n *Node) Insert(k int){
	switch {
	case k <= n.key && n.left == nil : 
		n.left = &Node{key: k, height: 1, depth: n.depth +1}
	case k > n.key  && n.right == nil :
		n.right = &Node{key: k, height: 1, depth: n.depth +1}
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
func(n *Node) SetDepth(){
	d := n.depth
	f := func(k *Node){
		if n.left != nil{
			n.left.height = d + 1
		}
		if n.right != nil{
			n.right.height = d + 1
		}
	}
	n.Traverse(f)	
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
	t.root.depth = 0
	d := func(n *Node) {
		if n.left != nil{
			n.left.depth = n.depth + 1 
		}
		if n.right != nil{
			n.right.depth = n.depth + 1
		}
	}
	t.root.Traverse(d)
}

func(t *Tree) left_rotate(){
	x := t.root.right
	t.root.left_rotate()
	t.root = x
	t.root.left.SetHeight()
	t.root.right.SetHeight()
}
func(t *Tree) right_rotate(){
	x := t.root.left
	t.root.right_rotate()
	t.root = x
	t.root.left.SetHeight()
	t.root.right.SetHeight()
}
// Left rotates move the previously right child node to being the 
// root node
//  1. Store the new and old root as r and o
//  2. Store the left child (l) of the new root, as if it is not nil it will need to be preserved

func(n *Node) left_rotate(){
	r := n.right
	o := n
	left := r.left
	r.left = o
	o.right = left
}
// Right Rotates move the previously left child node to being the
// root node
//  1. Store the new and old root as  and o
//  2. Store the right child(r) of the new root, as if it is not nil it will need to be preserved
func(n *Node) right_rotate(){ 
	r := n.left
	o := n
	right := r.right
	r.right = o
	o.left = right
}