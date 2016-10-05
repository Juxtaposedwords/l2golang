package main

import "fmt"
import "time"

func main() {

	i := 2
	fmt.Print("write ", i, " as ")
	switch i {             // Here's a basic switch
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// You can use commas to separate multiple expressions in the same case statement
  //		We use optional default case in this example as well.

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
}
