package main

import "fmt"

func main() {
	s := make([]string, 3)
//	Slices are typed only by the elements they contains (not the number of elements
//  To create an empty slice with non-zero length, use the builtin make. Here we make
//	a slice of strings of length 3 (intially zero-valued).
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

// len returns the length of the splice as expected
	fmt.Println("len: ", len(s))

// slices support several operations. One such is the builtin append, which returns a slice containing one or more new values
// Note that we have to accept a return valuef rom append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appd: ", s)


// Slices can also be copy'd. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

//  Slices support a `slice` operator with the syntax slice[low:high]
//   for example, this gets a slice of the elements s[2], s[3], and s[4]
	l := s[2:5]
	fmt.Println("sl1: ", l)

//  This slices up to (but excluding s[5]
	l = s[:5]
	fmt.Println("sl2: ", l)

// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3: ", l)


	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
// Slices can be composed of multi-dimensional data structures. 
// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
