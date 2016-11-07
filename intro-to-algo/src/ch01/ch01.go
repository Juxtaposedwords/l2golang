package ch01


func insertSort(A []int) {
		// store the number we're sorting on in 'key'
		key := A[j]
		// store the location in the array
		// This implicity tells us everything before i is sorted
		i := j -1
		fmt.Printf("Item: j: %d\n", j)
		fmt.Printf("%z\n", A)
		// From the last sort digit to the first (so Right to Left)
		//     if this item is greater than key, move that value one
		//     to the right and decrement the counter
		for i >=  0 && A[i] > key {
			A[i+1] = A[i]
			i = i -1
		} 
		A[i+1] = key

}
