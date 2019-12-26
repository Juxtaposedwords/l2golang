package fuel

func required(input uint64) uint64 {
	if input < 7 {
		return 0
	}
	return uint64(input/3) - 2
}
// TotalRequired computes the amount of fuel required given a list of weights.
func TotalRequired(inputs []uint64) uint64 {
	var sum uint64
	for _, entry := range inputs {
		sum += required(entry)
	}
	return sum
}
