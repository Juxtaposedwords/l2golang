package main

import "fmt"
import "math"

const s string = "constant"  //declaring a constant value

func main() {
	fmt.Println(s)        // The given constant can appear anywehre a var statement can
	
	const n = 500000000   // const statement can appear anywehre a var statement can
	
	const d = 3e20 / n    // const perform arithemtic with arbitrary precisions
	fmt.Println(d)        // A numermic constant has no type until it's given one, such as by explicit cast

	fmt.Println(int64(d)) // a numeric constant has no type until it's given one, such as by an explicity cast like int64()


	fmt.Println(math.Sin(n)) // A number can be given a type by using it in a context that requires one, such as a variable assignment or function call
													 // For example, here math.Sin expects a float64.
}
