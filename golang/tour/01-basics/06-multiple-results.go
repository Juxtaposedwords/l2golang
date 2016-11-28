package main

import "fmt"

func swap(x, y string) (string, string){
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	// note this is a classic unpacking where b == hello and a == world 
	// 		keep in mind swap return ['world', 'hello']
	fmt.Println(a, b)
}
