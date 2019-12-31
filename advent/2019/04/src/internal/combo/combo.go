package combo

import (
//"github.com/google/logger"
//	"io/ioutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sync"
	"sync/atomic"
)

/*Permutations uses a look ahead approach per digit to find all the possible combinations of possibilie between the two numbers in which all options in which no digit is descending and there are at least two repeated digits.
 */
func Permutations(lower, upper int) (int, error) {
	//defer logger.Init("LoggerExample", true, false, ioutil.Discard)

	switch {
	case digits(lower) != digits(upper):
		return 0, status.Errorf(codes.InvalidArgument, "different digit lengths: lower %d upper %d ", lower, upper)
	case lower > upper:
		return 0, status.Errorf(codes.InvalidArgument, "lower greater than upper: lower %d upper %d ", lower, upper)
	case digits(lower) == 1:
		return 0, nil
	}

	var combinations int
	for i := lower ; i <= upper; i++ {
		var repeats bool
		i, repeats = digitValidation(i)
		if repeats {
//		logger.Infof("i: %d\n", i)
			combinations++
		}
	}
	return combinations, nil
}

func digitValidation(input int) (int, bool) {
	digits := intToSlice(input)
	lowest := digits[0]
	observed := map[int]int{lowest:1}
	var valid bool
	for j := 1; j < len(digits); j++ {
		switch {
		case digits[j] < lowest:
			for k := j - 1; k < len(digits); k++ {
				digits[k] = lowest
			}
			return sliceToInt(digits) - 1, false
		case digits[j] == lowest:
		default:
			lowest = digits[j]
		}
		observed[digits[j]]++
	}
	for _, v := range observed{
		if v == 2 {
			valid=true
		}
	}
	return sliceToInt(digits), valid
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

// brute uses brute case to find all possible permutations. The function is not inteded for external use, but strictly for help in the creation of test cases.
func brute(lower, upper int) (int, error) {
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
			if validBrute(digit) {
				atomic.AddInt32(&result, 1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return int(result), nil
}
func validBrute(input int) bool {
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
