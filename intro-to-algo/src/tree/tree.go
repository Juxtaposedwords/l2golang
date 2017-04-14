package tree
import "fmt"
type Node struct{
	key int
	depth int
	left *Node
	right *Node
}

func(n *Node) Insert(k int){
	switch {
	case k <= n.key && n.left == nil : 
		n.left = &Node{key: k, depth: n.depth+1}
	case k > n.key  && n.right == nil :
		n.right = &Node{key: k, depth: n.depth+1}
	case k <= n.key  && n.left != nil : 
		n.left.Insert(k)
	case k > n.key  && n.right != nil :
		n.right.Insert(k)
	}
}
func(n *Node) Traverse() {
	fmt.Printf("%z\n",n)
	if n.left != nil {
		n.left.Traverse()
	}
	if n.right != nil {
		n.right.Traverse()
	}
}