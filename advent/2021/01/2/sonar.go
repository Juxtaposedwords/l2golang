package sonar

func Increased(depths []int) int {
	increased := 0
	for i := 1; i < len(depths); i++ {
		l, r := depths[i-1], depths[i]
		if l < r {
			increased++
		}
	}
	return increased
}

func WindowedIncreased(depths []int, window int) (int, error) {
	return 0, nil
}
