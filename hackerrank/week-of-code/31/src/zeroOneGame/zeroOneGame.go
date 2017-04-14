package zeroOneGame

import "sort"

func zeroOneGame(input []int) string {
	return ""
}

func validPlays(input []int) []int {
	output := []int{}
	for i, _ := range input {
		if i == 0 || i == len(input)-1 {
			continue
		}
		if input[i-1] == 0 && input[i+1] == 0 {
			output = append(output, i)
		}
	}
	return output
}
