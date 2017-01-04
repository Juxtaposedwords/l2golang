package linkedList
import (
	"testing"
//	"reflect"
)
func Test_delete( t *testing.T){
	a := CreateLinkedList(5)
	b := a.next
	b.delete()
	have := a.visitMap()
	visited := 0
	want := 4
	for k , v := range have {
		if v > 1{
			t.Errorf("Node %z was visited %d, not 1 time", k, v)
		}
		visited +=1
	}
	if visited != want {
		t.Errorf("Incorrect amount of nodes visited.\n VisitMap: %z\n",have)
	}
}
func Test_add( t *testing.T){
	a := CreateLinkedList(5)
	a.add()
	have := a.visitMap()
	want := 6
	visited := 0
	for k , v := range have {
		if v > 1{
			t.Errorf("Node %z was visited %d, not 1 time", k, v)
		}
		visited +=1
	}
	if visited != want {
		t.Errorf("You visisted %d when you should have visited", visited, want)
	}
}