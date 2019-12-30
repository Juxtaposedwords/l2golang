package combo

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
/*Brute uses brute force approach to find all the possible combinations of possibilie between the two numbers in which all options in which no digit is descending and there are at least two repeated digits.
*/
func Brute(lower, upper int) (int,error) {
	lowerBounds, upperBounds := sliceify(lower), sliceify(upper)

	switch {
	case len(lowerBounds) != len(upperBounds):
		return 0, status.Errorf(codes.InvalidArgument, "different digit lengths: lower %d upper %d ",lower, upper)
	case lower > upper:
		return 0, status.Errorf(codes.InvalidArgument, "lower greater than upper: lower %d upper %d ",lower, upper)
	}
	var result int
	for i := lower; i <= upper; i ++ {
		if valid(i) {
			result++
		}
	}
return result, nil
}

 func digits(i int) (count int) {
 	for i != 0 {
 		i /= 10
 		count = count + 1
 	}
 	return count
 }
func valid(input int) bool {
	var intSlice []int
	intSlice = sliceify(input)
	if len(intSlice) == 1 {
		return false
	}
	var repeat bool
	for i := 1; i < len(intSlice); i++ {
		switch {
		case intSlice[i-1] > intSlice[i]: 	
			return false
		case intSlice[i-1] == intSlice[i] :
			repeat = true
		}
	}
	if !repeat {
		return false
	}
	return true

}
func recurSliceify(input int, output []int) []int {
    if input != 0 {
        i := input % 10
        output = append([]int{i}, output...)
        return recurSliceify(input/10, output)
    }
    return output
}
func sliceify( input int) []int {
	var output []int
	return recurSliceify(input, output)
}