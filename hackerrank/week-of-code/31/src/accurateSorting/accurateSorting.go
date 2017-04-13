package accurateSorting

const (
	hrTrue  = "Yes"
	hrFalse = "No"
)

func IsAbsSortable(length int, unsorted []int) string {
	// start with the first item
	highest := unsorted[0]
	for _, entry := range unsorted {
		// if we ever find an absolute value > 1 one it will be unswappable.
		//    the array is beyond saving
		if highest-entry > 1 {
			return hrFalse
		}

		// We have to set our "highest" point, to show
		if highest < entry {
			highest = entry
		}
	}
	return hrTrue
}
