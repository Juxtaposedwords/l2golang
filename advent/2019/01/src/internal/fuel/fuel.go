package fuel

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"strconv"
	"bufio"
	"os"
)

const maxUINT64 = ^uint64(0)

func required(input uint64) uint64 {
	if input < 7 {
		return 0
	}
	requiredFuel := uint64(input/3) - 2
	return requiredFuel + required(requiredFuel) 
}

func totalRequired(inputs []uint64) uint64 {
	var sum uint64
	for _, entry := range inputs {
		sum += required(entry)
	}
	return sum
}

// ReadFuel takes in the path to a file which contains all the weights you have.
//  The amount of fuel necessary for each item is computed and added to a running 
//  sum which is returned at the end.  
func ReadFuel(filepath string) (uint64, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return 0, status.Error(codes.FailedPrecondition, err.Error())
	}
	return readInts(bufio.NewReader(f))

}
func readInts(r io.Reader) (uint64, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result uint64
	for scanner.Scan() {
		x, err := strconv.ParseUint(scanner.Text(), 10,64)
		if err != nil {
			return result, status.Error(codes.InvalidArgument, err.Error())
		}
		diff := required(x)
	//	logger.Infof("%d %d",maxUINT64-result,x)
		if maxUINT64  - result < diff {
			return 0, status.Error(codes.Internal, "integer overflow")
		}

		result +=diff 
	}
	return result, scanner.Err()
}
