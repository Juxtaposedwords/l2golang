package main

import (
	"github.com/google/logger"
	"internal/geo"

)
func main() {
	left := []string{}
	right := []string{}
	resp, err := geo.QuickestIntersection(left, right)
	if err != nil  {
		logger.Fatalf("error: %#v", err)
	}
	logger.Infof("Answer: %d", resp)
}