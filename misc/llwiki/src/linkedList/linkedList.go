package linkedList
import (
	"fmt"
	"math/rand"
)

type Node struct{
	Label string
	Next *Node
}
type NodeList []*Node

func CreateLinkedList(c int) (*Node){
	var first, last, swap *Node
	first = &Node{}
	last = first
	for i:=2; i <= c; i++{
		swap = &Node{}
		last.Next = swap
		last = swap
	}
	swap.Next = first
	return first
}

func (h *Node)toNodeList()(NodeList){
	var result []*Node
	f := func(n *Node){
		result = append(result, n)
	}
	h.traverse(f)
	return result
}

func (n *Node)add(){
	nn := &Node{Next: n.Next}
	n.Next = nn
}

func (h *Node)delete(){
	var prev *Node 
	f := func(n *Node) {
		if n.Next == h {
			prev = n
		}
	}
	h.traverse(f)
	prev.Next = h.Next
}

func (h *Node)shuffle(){
	s := h.toNodeList()
	s.shuffle()
}
func (h *Node)listLength()(int){
	i := 0
	f := func(k *Node){
		i+=1
	}
	h.traverse(f)
	return i
}

func (h *Node)visitMap()(map[*Node]int){
	r := make(map[*Node]int)
	f := func(n *Node){
		r[n] +=1 
	}
	h.traverse(f)
	return r
}
func (h *Node)traverse(f func(*Node)){
	curr := h
	for {
		f(curr)
		curr = curr.Next
		if curr == h{
			break
		}
	}
}
func (l NodeList)shuffle(){
	for i := 0; i <len(l) - 2; i++ {
		j  := rand.Intn(len(l) - i - 1)
		l[i], l[i+j] = l[i+j], l[i]
	}
	for i := range l {
		l[i].Next = l[(i + 1 % len(l))]
	}
}
func (h *Node)printer(){
	h.traverse(func(n *Node){fmt.Printf("Label: %z\n", n.Label)})
}
