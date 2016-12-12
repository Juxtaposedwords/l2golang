package doublyLinkedList
import (
	"fmt"
	"math/rand"
)

type node struct{
	label int
	next *node
	prev *node
}

type nodeList []*node

func CreateLinkedList(c int) (*node){
	var first, last, swap *node
	first = &node{label: 1}
	last = first
	for i:=2; i<= c; i++{
		swap = &node{label:i, prev: last}
		last.next = swap
		last = swap
	}
	swap.next = first
	first.prev = swap
	return first
}

func (n *node)add(a int){
	nn := &node{label: a, prev: n, next: n.next}
	n.next.prev = nn
	n.next = nn
}
func (n *node)delete(){
	n.prev.next = n.next
	n.next.prev = n.prev
}
func (n *node)listLength()(a int){
	i :=0
	f := func(k *node){
		i +=1
	}
	n.traverse(f)
	return i
}

func (h *node)toNodeList() (nodeList){
	var result []*node
	f := func(n *node){
		result = append(result, n)
	}
	h.traverse(f)
	return result
}

func (h *node)shuffle(){
	s := h.toNodeList()
	s.shuffle()
}

func (h *node)printer(){
	h.traverse(func(n *node){fmt.Printf("label: %z\n", n.label)})
}

func (h *node)visitMap()(map[int]bool){
	r := make(map[int]bool)
	f := func(n *node){
		r[n.label] = true
	}
	h.traverse(f)
	return r
}
// attempt to traverse the entire linked list until you hit the starting node
//  so take not that this can be easily broken by making an incomplete cycle
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
// Attempt to walk through the cycle on the raverse methodology
func (h *node)reverseTraverse(f func(*node)) {
	curr := h
	for {
		f(curr)
		curr = curr.prev
		if curr == h {
			break
		}
	}
}

// shuffles the order of the nodelist
func (s nodeList)shuffle(){
	for i := 0;i < len(s) -2; i ++{
		j := rand.Intn(len(s) -i - 1)
		s[i], s[i+j] = s[i+j], s[i]
	}
	for i := range s {
		s[i].next = s[(i+1) % len(s)]
	}
}