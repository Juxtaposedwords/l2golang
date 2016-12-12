package llShuffle
import (
	"fmt"
	"math/rand"
)
type node struct{
	label int
	next *node
	prev *node
}
// Created a class so we add some methods specifically
type nodeList []*node

// Creates a circular linked list of a nodes 
func CreateLinkedList(a int) (*node){
	var first,last, swap *node
	first = &node{label: 1}
	last = first
	for i:=2; i <= a; i++{
		swap = &node{label: i}
		last.next = swap
		last = swap
	}
	swap.next = first
	return first
}
func CreateDoublyLinkedList(a int) (*node){
	var first,last, swap *node
	first = &node{label: 1}
	last = first
	for i:=2; i <= a; i++{
		swap = &node{label: i, prev: last}
		last.next = swap
		last = swap
	}
	swap.next = first
	first.prev = swap
	return first
}

//Turns our circular linked list into an array of nodes
func (h *node)toNodeList() (nodeList){
	var result []*node
	f := func(n *node) {
		result = append(result, n)
	}
	h.traverseList( f)
	return result
}

func (h *node)visitMap()(map[int]bool){
	r := make(map[int]bool)
	f := func(n *node){
		r[n.label] = true
	}
	h.traverseList(f)
	return r
}

//Shuffle the linked list from just the head node
//	Makes a list of the cycle, then shuffles the nodes
func (head *node)shuffle(){
	s  := head.toNodeList()
	s.shuffle()
}

func (n *node)delete(){
	var prev *node
	f := func(k *node) {
			if k.next == n {
				prev = k
			}
		}
	n.traverseList(f)
	prev.next = n.next
}
func (n *node)add(a int){
	var nn *node
	if n.prev != nil{
		nn =  &node{label: a, prev: n, next: n.next}		
		n.next.prev = nn		
	} else {
		nn =  &node{label: a, next: n.next}		
	}
	n.next = nn
}

//Shuffles the order of the nodeList
//   Done only to the array form, as it's just easier
func (s nodeList)shuffle(){
	for i :=0; i < len(s) -2; i++{
		j := rand.Intn(len(s) - i - 1)
		s[i], s[i+j] = s[i+j], s[i]
	}
	for i := range s {
		s[i].next = s[(i+1) % len(s)]
	}
}

// Print out the labels from the NodeList
func (s nodeList)labels()([]int){
	var a []int
	m := func (n *node) { 
		a = append(a, n.label) }
	s[0].traverseList( m)
	return a
}

// Print out the entire nodeList
func (h *node)printer() {
	h.traverseList(func(n *node) { fmt.Printf("%d\n", n.label)})
	fmt.Println()
}

// Mapping function for traversing through the linked lists
func (h *node)traverseList(f func(*node)){
	curr := h
	for {
		f(curr)
		curr = curr.next
		if curr == h {
			break
		}
	}

}
func (h *node)reverseTraverseList(f func(*node)){
	if h.prev == nil{
		return
	}
	curr := h
	for {
		f(curr)
		curr = curr.prev
		if curr == h {
			break
		}
	}
}