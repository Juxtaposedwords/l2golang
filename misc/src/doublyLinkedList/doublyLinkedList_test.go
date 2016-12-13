package doublyLinkedList
import (
	"testing"
//	"reflect"
)
func Test_delete( t *testing.T){
	a := CreateLinkedList(5)	
	b := a.next
	a.delete()
	have := b.visitMap()
	want := 4
	visited := 0
	for k, v := range have{
		if v > 1 {
			t.Errorf("Node %z was visited %d times, not 1 time", k, v)
		}
		visited +=1
	}
	if visited != want {
		t.Errorf("You visited %d nodes when you should have visited %d nodes", visited, want)

	}
}
func Test_add( t *testing.T){
	a := CreateLinkedList(5)
	a.add()
	have := a.visitMap()
	want := 6
	visited :=0
	for k,v := range have{
		if v > 1 {
			t.Errorf("Node %z was visisted %d times, not 1 time", k,v )
		}
		visited +=1
	}
	if visited != want{
		t.Errorf("The list has %d unique nodes, not %d", visited,want)
	}
}

func Test_shuffle( t *testing.T){
	//Stubbing this out for later work
}