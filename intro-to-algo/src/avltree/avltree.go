package avltree
import (
	"fmt"
)
type Node struct{
	key int
	depth int
	height int
	left *Node
	right *Node
}
type Tree struct{
	root *Node
}

func(t *Tree) Insert(k int){
	if t.root == nil{
		t.root = &Node{key: k, depth: 0, height: 0}
	} else {
		t.root.Insert(k)
	}
}
func(n *Node) Insert(k int){
	switch {
	case k <= n.key && n.left == nil : 
		n.left = &Node{key: k, depth: n.depth+1, height: 0}
	case k > n.key  && n.right == nil :
		n.right = &Node{key: k, depth: n.depth+1, height: 0}
	case k <= n.key  && n.left != nil : 
		n.left.Insert(k)
	case k > n.key  && n.right != nil :
		n.right.Insert(k)
	}

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
	fmt.Printf("%z\n", n)
	if n.left != nil {
		n.left.Traverse()
	}
	if n.right != nil {
		n.right.Traverse()
	}
}

func left_rotate(r *Node)(*Node){
	var left,p *Node  
	p = r.right
	if p.left != nil {
		left = p.left
	}
	p.left = r
	r.right = left
	return p
}
func right_rotate(r *Node)(*Node){
	var right,p *Node  
	p = r.left
	if p.right != nil {
		right = p.right
	}
	p.right = r
	r.left = right
	return p
}