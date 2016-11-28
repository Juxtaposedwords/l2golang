package main
import "fmt"
func main() {
//	To create an empty map, use the builtin make:
//		make(map[key-type] val-type)
	m := make(map[string]int)

//	set key/value pairs when using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

// 	Printing a map with e.g. Println will show all of its key/value pairs
	fmt.Println("map: ", len(m))

//	get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1:  ", v1)

// the builtin len returns the number of key/value pairs when on a map
	fmt.Println("len: ", len(m))


// the builtin delete removes key/values pairs form a map
	delete(m, "k2")
	fmt.Println("map: ", m)

// The optional second return value when getting a value from a map indicates if hte key was present in the map.
//		This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn't
//		need all the value itself, so we ignored it if with the blank identifier _ 
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

//	you can also delcare and intiailize a new map in the same line with this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

}
