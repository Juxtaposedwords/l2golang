package avltree
import (
	"fmt"
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
		t.root = &Node{key: k, height: 0, depth: 0}
	} else {
		t.root.Insert(k, t)
		//t.Rebalance()
	}
}
func(n *Node) Insert(k int, t *Tree){
	switch {
	case k < n.key && n.left == nil : 
		n.left = &Node{key: k, height: 1, depth: n.depth +1}
	case k >= n.key  && n.right == nil :
		n.right = &Node{key: k, height: 1, depth: n.depth +1}
	case k < n.key  && n.left != nil : 
		n.left.Insert(k,t)
	case k >= n.key  && n.right != nil :
		n.right.Insert(k,t )
	}
	n.SetHeight()
	n.Rebalance(t)
}
// A node's height is the greater of it's two children + 1
func(n *Node) SetHeight(){
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}

	if left < right {
		n.height = right + 1
	} else if left >= right {
		n.height = left + 1
	}	
}

// A node's depth is the shortest number of connections to the root
func(n *Node) SetDepth(){
	f := func(k *Node){
		if n.left != nil{
			n.left.depth = n.depth + 1
		}
		if n.right != nil{
			n.right.depth = n.depth + 1
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
	t.root = t.root.Rebalance(t)
}
func(n *Node) Rebalance(t *Tree) (*Node){
	left, right := 0,0
	if n.left != nil {
		left = n.left.height
	} 
	if n.right != nil{
		right = n.right.height 
	}
	diff := left - right
	a := t.root
	switch {
	case diff < -1:
		a = n.right
		n.left_rotate()
	case diff > 1: 
		a = n.left
		n.right_rotate()
	}	
	n.SetHeight()
	n.SetDepth()
	fmt.Printf("Returning %d at node %d\n",n.key,a.key)
	return a
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