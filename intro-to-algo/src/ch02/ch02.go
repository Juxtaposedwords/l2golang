package ch02

import "log"

func insertSort(A []int) {
	for j := 0; j < len(A); j++ {
		// store the number we're sorting on in 'key'
		key := A[j]
		// store the location in the array
		// This implicity tells us everything before i is sorted
		i := j -1
		// From the last sort digit to the first (so Right to Left)
		//     if this item is greater than key, move that value one
		//     to the right and decrement the counter
		for i >=  0 && A[i] > key {
			A[i+1] = A[i]
			i = i -1
		}
		A[i+1] = key
	}
}

func reverseInsertSort(A []int){
	for j:= len(A)-1; j >= 0; j-- {
		// store the number we're sorting on in 'key'
		key := A[j]
		// store the location in the array
		// This implicity tells us everything before i is sorted
		i := j + 1
		//max := len(A)
		// From the last sort digit to the first (so left to right)
		//     if this item is greater than key, move that value one
		//     to the right and decrement the counter
		for i < len(A) && A[i] < key {
			A[i-1] = A[i]
			i = i + 1
		}
		A[i-1] = key
	}
}
func mergeSort (a []int) []int{
	switch len(a) {
	case 0, 1:
		return a
	case 2:
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	default:
		b, c := split(a)
		b = mergeSort(b)
		c = mergeSort(c)
		return merge(b,c)
	}
}

func merge( a, b []int) []int {
	var sorted []int
	for {
		switch {
		case len(a) == 0 && len(b) == 0:
			return sorted
		case len(a) == 0:
			sorted = append(sorted, b...)
			return sorted
		case len(b) == 0:
			sorted = append(sorted, a...)
			return sorted
		case a[0] > b[0]:
			sorted = append(sorted, b[0])
			b = b[1:]
		case b[0] >= a[0]:
			sorted = append(sorted, a[0])
			a = a[1:]
		default:
			log.Fatal("This should never happen.")
		}
	}
	return sorted
}
func split( a []int) ([]int, []int){
	x := len(a)/2
	b := a[x:]
	a = a[:x]
	return a, b
}
