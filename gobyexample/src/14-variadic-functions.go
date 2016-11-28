package main

import "fmt"
// vardiac functions can be called with any number of trailing arguments. For example,
//  fmt.Println is a common variadic function

// Here's a function that will take an aribtrary number of ints as arguments
func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {
// Variadic functions can be called in the usual way with individual arguments
  sum(1, 2)
  sum(1, 2, 3)

  nums := []int{1, 2, 3, 4}

//  If you have multiple args in a slice, apply them to a variadic function using func(slice...) like this. If you have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
  sum(nums...)
}
