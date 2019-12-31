package main

import (
	"internal/intcode"
	"github.com/google/logger"

)
var inputSlice = []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,19,5,23,2,23,13,27,1,10,27,31,2,31,6,35,1,5,35,39,1,39,10,43,2,9,43,47,1,47,5,51,2,51,9,55,1,13,55,59,1,13,59,63,1,6,63,67,2,13,67,71,1,10,71,75,2,13,75,79,1,5,79,83,2,83,9,87,2,87,13,91,1,91,5,95,2,9,95,99,1,99,5,103,1,2,103,107,1,10,107,0,99,2,14,0,0}
func main() {
	test :=inputSlice
	noun, verb, err :=intcode.BrutePair(test,19690720)
	if err != nil {
		logger.Fatalf("error generating list: %#v",err)
	}
	logger.Infof("noun: %d verb: %d \n Answer: %d\n",noun, verb, (noun*100)+verb)
}
