package llShuffle

import (
	"testing"
	"reflect"
)

func Test_shuffle(t *testing.T){
	head := CreateLinkedList(50)
	original := head.toNodeList()
	have := original.labels()
	original.shuffle()
	want := original.labels()
	if reflect.DeepEqual(have, want) {
		t.Errorf("No shuffle,\n  have: %z\n  want: %z\n", have,want)
	}
}

func Test_add_ll_items(t *testing.T){
	x := CreateLinkedList(4)
	a := x.toNodeList().labels()
	for want, have := range(a){
		// offset the 0 based array
		want +=1
		if want != have {
			t.Errorf("have: %d , want: %d . MisMatch on generation.\n",  have, want)
		}
	}
	want := x.visitMap()
	x.shuffle()
	have := x.visitMap()
	if ! reflect.DeepEqual(want, have) {
		t.Errorf("have: %d\nwant: %d\n Description: \n", want, have)
	}

	//x.delete()
	//head := CreateLinkedList(3)
}
func Test_delete_ll_items(t *testing.T){
	a := CreateLinkedList(5)
	want := a.visitMap()
	delete(want, 1)
	b := a.next
	a.delete()
	have := b.visitMap()
	desc := "Delete was unsucessful"
	if ! reflect.DeepEqual(want, have) {
		t.Errorf("have: %d\nwant: %d\n Description: %s \n", want, have, desc)
	}

	//x.delete()
	//head := CreateLinkedList(3)
}

func Test_edc_shuffle(t *testing.T) {
/*	tcs := []struct{
		have int
		want []node
		desc string
		}{ 
			{25, }
		}
	}*/
	x := CreateDoublyLinkedList(50)
	x.next.delete()
	x.add(2)

	//x.prev.reverseTraverseList(func(n *node) { fmt.Printf("%d\n", n.label)})
}
