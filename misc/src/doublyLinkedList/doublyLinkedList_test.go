package doublyLinkedList
import (
	"testing"
	"reflect"
)
func Test_delete( t *testing.T){
	a := CreateLinkedList(5)	
	want := a.visitMap()
	delete(want, 1)
	next  := a.next
	a.delete()
	have := next.visitMap()
	desc := "Delete was unsuccessful"
	if ! reflect.DeepEqual(want, have){
		t.Errorf("\nhave: %z\nwant: %z\nDescription: %s\n",have, want, desc)
	}
}
func Test_add( t *testing.T){
	a := CreateLinkedList(5)
	want := 6
	a.add(9)
	have := a.listLength()
	desc := "Add was unsuccessful"
	if ! reflect.DeepEqual(want, have){
		t.Errorf("have: %z   want: %z    Description: %s\n",have, want, desc)
	}
}

func Test_shuffle( t *testing.T){
	//Stubbing this out for later work
}