package main

// for is Go's only looping construct. Here are three basics types of for loops.


import "fmt"

func main() {
	i := 1

	// The most basic type, with a single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic intial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}


	// For without a condition will loop repeatedly until you break out of the loop
	//	or return from the enclosing function
	for{
		fmt.Println("loop")
		break
	 }

}
