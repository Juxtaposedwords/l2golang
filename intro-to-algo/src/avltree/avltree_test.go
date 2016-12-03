package avltree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	 tcs := []struct{
		have []int
		want []int
		desc string
	}{
  		{[]int{1, 7, 2, 3, 4, 2, 3, 4, 5}, []int{1, 2, 2, 3, 3, 4, 4, 5, 7}, "typical case"},
  	//	{[]int{5, 4, 3, 3, 2, 2, 1}, []int{1, 2, 2, 3, 3, 4, 5}, "reversed"},
	}

	for _, test := range(tcs){
		tree := &Tree{} 
		for _, i := range(test.have){
			tree.Insert(i)
		}
		tree.root.Traverse()
	}


}