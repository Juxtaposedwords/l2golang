package tree

type Node struct{
	key int
	left *Node
	right *Node
}

func(n *Node) Insert(k int){
	switch {
	case k <= n.key && n.left == nil : 
		n.left = &Node{key: k}
	case k > n.key  && n.right == nil :
		n.right = &Node{key: k}
	case k <= n.key  && n.left != nil : 
		n.left.Insert(k)
	case k > n.key  && n.right != nil :
		n.right.Insert(k)
	}
}