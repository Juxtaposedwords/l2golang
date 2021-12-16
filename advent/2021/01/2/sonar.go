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

func WindowedIncreased(depths []int) int {
	windowDepths := []int{}
	for i := 2; i < len(depths); i++ {
		depth := depths[i-2] + depths[i-1] + depths[i]
		windowDepths = append(windowDepths, depth)
	}
	return Increased(windowDepths)
}
