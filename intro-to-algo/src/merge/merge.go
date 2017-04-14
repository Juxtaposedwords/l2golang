package merge

import "log"

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
