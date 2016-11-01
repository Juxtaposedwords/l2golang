// Generate K unique integers from 1 - 27,0100
package main
import "fmt"

// none of this works yo
// Given:
//  1. Knuth's Birthday Problem
//  2. The cheapness of resources
// I should generate 27000 int array and remove at a random

func makeNumber(max int) {
  collection := make([]int, max)
  for i := 1; i <= max; i++{
    collection[i-1] = i
  }
  return collection
}


func main() {
  collection := makeNumber(27000)
  fmt.Println(collection)
}
