package ch01


func insertSort(A []int) {
	for j := 0; j < len(A); j++ {
		key := A[j]
		i := j -1
		for i >=  0 && A[i] > key {
			A[i+1] = A[i]
			i = i -1
		} 
		A[i+1] = key
	}

}
