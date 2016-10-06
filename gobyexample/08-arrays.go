package main

import "fmt"

func main() {
	var a [5]int // we create an array that will hold exactly 5 ints
							 // The type of elements and length are both part of the array's type.
							 // By default an array is zero-valued, which for ints means 0s.
	fmt.Println("emp:", a)

	a[4] = 100  // We can set a value at an index using the array[index] = value syntac, and get the value.
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))  //the builtin len returns the length of the array

	b := [5]int{1, 2, 3, 4, 5}		//Use this syntax to declare and intialize an array in one line
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
