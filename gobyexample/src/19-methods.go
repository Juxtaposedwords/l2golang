package main

import "fmt"

type rect struct {
  width, height int
}

// This area method has a receiver type of *rect
func (r *rect) area() int {
  return r.width * r.height
}

// Methods can be defined for either pointer or value receiever types.
//  Here's an example of a value receiver
func (r rect) perim() int {
  return 2*r.width +  2 * r.height
}

func main () {
  r := rect{width: 10, height: 5}

// Here we call 3 methods defined on our struct.
  fmt.Println("area:  ", r.area())
  fmt.Println("perim: ", r.perim())


// Golang automatically handles conversion ebtwen values and pointesr for method calls.
//    You may want to avoid copying on method calls or allow hte method to mutate the receiving struct
  rp := &r
  fmt.Println("area:  ", rp.area())
  fmt.Println("perim: ", rp.perim())
}
