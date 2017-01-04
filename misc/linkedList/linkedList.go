package linkedList
import (
	"fmt"
	"math/rand"
)

type node struct{
	label string
	next *node
}
type nodeList []*node

func CreateLinkedList(c int) (*node){
	var first, last, swap *node
	first = &node{}
	last = first
	for i:=2; i <= c; i++{
		swap = &node{}
		last.next = swap
		last = swap
	}
	swap.next = first
	return first
}

func (h *node)toNodeList()(nodeList){
	var result []*node
	f := func(n *node){
		result = append(result, n)
	}
	h.traverse(f)
	return result
}

func (n *node)add(){
	nn := &node{next: n.next}
	n.next = nn
}

func (h *node)delete(){
	var prev *node 
	f := func(n *node) {
		if n.next == h {
			prev = n
		}
	}
	h.traverse(f)
	prev.next = h.next
}

func (h *node)shuffle(){
	s := h.toNodeList()
	s.shuffle()
}
func (h *node)listLength()(int){
	i := 0
	f := func(k *node){
		i+=1
	}
	h.traverse(f)
	return i
}

func (h *node)visitMap()(map[*node]int){
	r := make(map[*node]int)
	f := func(n *node){
		r[n] +=1 
	}
	h.traverse(f)
	return r
}
func (h *node)traverse(f func(*node)){
	curr := h
	for {
		f(curr)
		curr = curr.next
		if curr == h{
			break
		}
	}
}
func (l nodeList)shuffle(){
	for i := 0; i <len(l) - 2; i++ {
		j  := rand.Intn(len(l) - i - 1)
		l[i], l[i+j] = l[i+j], l[i]
	}
	for i := range l {
		l[i].next = l[(i + 1 % len(l))]
	}
}
func (h *node)printer(){
	h.traverse(func(n *node){fmt.Printf("label: %z\n", n.label)})
}
