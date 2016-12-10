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
/*func Test_edc_shuffle(t *testing.T) {
	tcs := []struct{
		have int
		want []node
		desc string
		}{ 
			{25, }
		}
	}
	x := CreateLinkedList(50)
	x.shuffle()
	x.printer()
}*/
