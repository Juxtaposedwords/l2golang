package main

import(
	"internal/fuel"
	"github.com/google/logger"

)
var (
	filepath = "/workspace/l2golang/advent/2019/01/src/input.txt"
)
func main() {
	resp, err := fuel.ReadFuel(filepath)
	if err != nil {
		logger.Fatalf("%#v",err)
	}
	logger.Infof("Answer is: %d", resp)
}