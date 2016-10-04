package main

import "fmt"
import "math"

const s string = "constant"  //declaring a constant value

func main() {
	fmt.Println(s)  // The given constant can appear anywehre a var statement can
	
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))


	fmt.Println(math.Sin(n))
}
