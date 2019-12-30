package combo

import (
	"github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sync"
	"sync/atomic"
)

/*Brute uses brute force approach to find all the possible combinations of possibilie between the two numbers in which all options in which no digit is descending and there are at least two repeated digits.
 */
func Brute(lower, upper int) (int, error) {
	switch {
	case digits(lower) != digits(upper):
		return 0, status.Errorf(codes.InvalidArgument, "different digit lengths: lower %d upper %d ", lower, upper)
	case lower > upper:
		return 0, status.Errorf(codes.InvalidArgument, "lower greater than upper: lower %d upper %d ", lower, upper)
	}
	var result int32
	var wg sync.WaitGroup
	for i := lower; i <= upper; i++ {

		wg.Add(1)
		go func(digit int) {
			if valid(digit) {
				atomic.AddInt32(&result, 1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return int(result), nil
}

/*Permutations uses a recursive programming solution to find all possible solutions between the two provided bounds.
 */
func Permutations(lower, upper int) (int, error) {
	switch {
	case digits(lower) != digits(upper):
		return 0, status.Errorf(codes.InvalidArgument, "different digit lengths: lower %d upper %d ", lower, upper)
	case lower > upper:
		return 0, status.Errorf(codes.InvalidArgument, "lower greater than upper: lower %d upper %d ", lower, upper)
	case digits(lower) == 1:
		return 0, nil
	}
	
	var combinations int
	for i := lower+1; i < upper; i++ {
		logger.Infof("i: %d\n", i)
		var repeat bool
		i, repeat = valid2(intToSlice(i))
		if repeat {
			combinations++
		}
	}
	return combinations, nil
}

func valid2(digits []int) (int, bool) {
	lowest := digits[0]
	var repeat bool
	for j := 1; j < len(digits); j++ {
		switch {
		case digits[j] < lowest:
			for k := j - 1; k < len(digits); k++ {
				digits[k] = lowest
			}
			return sliceToInt(digits) - 1, false
		case digits[j] == lowest:
			repeat = true
		default:
			lowest = digits[j]
		}
	}
	return sliceToInt(digits), repeat
}
func setter(index int, input []int) int {
	for i := index; index < len(input); i++ {
		input[i] = input[index]
	}
	return sliceToInt(input)

}
func digits(i int) (count int) {
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}
func valid(input int) bool {
	var intSlice []int
	intSlice = intToSlice(input)
	if len(intSlice) == 1 {
		return false
	}
	var repeat bool
	for i := 1; i < len(intSlice); i++ {
		switch {
		case intSlice[i-1] > intSlice[i]:
			return false
		case intSlice[i-1] == intSlice[i]:
			repeat = true
		}
	}
	if !repeat {
		return false
	}
	return true

}
func recurIntToSlice(input int, output []int) []int {
	if input != 0 {
		i := input % 10
		output = append([]int{i}, output...)
		return recurIntToSlice(input/10, output)
	}
	return output
}
func intToSlice(input int) []int {
	var output []int
	return recurIntToSlice(input, output)
}
