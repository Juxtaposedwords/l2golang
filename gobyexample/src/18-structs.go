package main

import "fmt"

// This person struct type has name and age fields
type person struct {
  name string
  name int
}



func main() {
  fmt.Println(person{"Bob", 20})                     // This syntax creates a new struct

  fmt.Println(person{name: "Alice", age: 30})        // You can anme the fields when intializing a struct

  fmt.Println(person{name: "Fred"})                  // Omitted fields willbe zero-valued

  fmt.Println(&person{name: "Ann", age: 40})         // An & prefix yeilds a pointer to the struct

  s := person{name: "Sean", age: 50}                 // Access struct fields with a dot
  fmt.Println(s.name)

  sp := &s                                           // You can also use dots with struct pointesr - the pointers are automatically deferenced
  fmt.Println(sp.age)

  sp.age = 51                                        // Structs are mutable
  fmt.Println(sp.age)
}
