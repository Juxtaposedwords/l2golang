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
func(n *Node) Traverse()  {
	if n.left != nil {
		n.left.Traverse()
	}
	// your action here
	if n.right != nil {
		n.right.Traverse()
	}
}
func(n *Node) Traverse2(f func(*Node))  {
	if n.left != nil {
		n.left.Traverse2(f)
	}
	f(n)
	if n.right != nil {
		n.right.Traverse2(f)
	}
}
func(t *Tree) Traverse() {
	t.root.Traverse()
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

func(t *Tree) left_rotate(){
	n := t.root.right
	o := t.root
	left := n.left
	n.left = o
	o.right = left
	t.root = n
	t.root.height = 0
	o.SetHeight()
}
func(t *Tree) right_rotate(){ 
	n := t.root.left
	o := t.root
	right := n.right
	n.right = o
	o.left = right
	t.root = n
	t.root.height = 0
}