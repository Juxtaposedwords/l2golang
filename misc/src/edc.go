package edc

type Node struct{
	label int
	next *Node
}
func CreateList(a int) (*Node){
	var first,last, swap *Node
	first = &Node{label: 1}
	last = first
	for i:=2; i <= a; i++{
		swap = &Node{label: i}
		last.next = swap
		last = swap
	}
	swap.next = first
	return first
}
