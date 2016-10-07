package main

// Now we'll show how pointers work in contrast to values with 2 functions: zeroval and zeoptr.
// zeroval - has an int parameter, so arguments will be baseed to it by value. It will get a copy of ival distinct from the 
//           one in the calling function
//  zeroptr - in contrast it has an *int parameter, meaning that it takes an int pointer. the *iptr code in the function body
//            then dereferences the pointer from its memory address to the current value at the address. Assigning a value to a
//            dereferenced pointer changes the value at the referenced address. 
import "fmt"


func zeroval(ival int) {
  ival = 0
}

func zeroptr(iptr *int) {
  *iptr = 0
}


func main() {
  i := 1
  fmt.Println("initial: ", i)

  zeroval(i)
  fmt.Println("zeroval: ", i)

//  The &i syntax gives the memory address of i, i.e. a pointer to i
  zeroptr(&i)
  fmt.Println("zeroptr: ", i)

// pointers can be printed too
  fmt.Println("pointer: ", &i)
}

// zeroval doesn't change the i in main, but zeroptr does becuase it has reference to the memory address for that variable
