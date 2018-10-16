package best

import ()

func maxProfit(prices []int) int {
	var max int
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
