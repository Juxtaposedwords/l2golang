package main

import (
	"fmt"
)


func main() {
	var t, i uint
	t, i = 1, 1
	for i = 1 ; i < 10 ; i++ {
		var w uint
		w = 3
		w |= t<<i
        fmt.Printf("%d << %d = %d \n", t , i , t<<i)
        fmt.Printf("%d\n",w)
    }
}