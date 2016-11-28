package main

import "fmt"

// a function can take zer or more argumetns
// here `add` takes two int parametersA
// Notice the type comes AFTER the variable
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42,13))
}
