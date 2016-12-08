package edc

import (
	"fmt"
	"testing"
)


func Test_edc_shuffle(t *testing.T) {
	a := 1
	x :=CreateList(50)
	starting := x.label
	x = x.next
	for x.label != starting {
		a+=1
		x = x.next
	}
	fmt.Printf("%d\n",a)

}