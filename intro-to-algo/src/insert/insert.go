package insert


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
