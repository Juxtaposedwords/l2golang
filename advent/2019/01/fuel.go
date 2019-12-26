package fuel

func required(input uint64) uint64 {
	return uint64(input/3) - 2
}
func TotalRequired(inputs []uint64) uint64 {
	var sum uint64
	for _, entry := range inputs {
		sum += required(entry)
	}
	return sum
}
