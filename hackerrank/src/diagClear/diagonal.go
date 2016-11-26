package diagClear

import (
	"fmt"
	"io"
)

func diagValue(r io.Reader) (int, error) {
	a, err := readInt(r)
	if err != nil {
		return 0, err
	}
	b, err := loadMatrix(a, r)
	if err != nil {
		return 0, err
	}
	c := sumSlice(primaryDiag(b))
	d := sumSlice(secondaryDiag(b))
	e := Abs(c - d)
	return e, nil
}

func sumSlice(a []int) int {
	var sum int
	for _, i := range a {
		sum += i
	}
	return sum
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func primaryDiag(x [][]int) []int {
	a := len(x[0]) - 1
	b := []int{}
	for i := 0; i <= a; i++ {
		b = append(b, x[i][i])
	}
	return b
}
func secondaryDiag(x [][]int) []int {
	a := len(x[0])
	b := a
	c := []int{}
	for i := 0; i < b; i++ {
		a -= 1
		c = append(c, x[i][a])
	}
	return c
}
func loadMatrix(x int, r io.Reader) ([][]int, error) {
	a := [][]int{}
	for i := 0; i < x; i++ {
		b, err := loadSlice(x, r)
		if err != nil {
			return nil, err
		}
		if len(b) != x {
			err := fmt.Errorf("The %dth row is not %d long", i+1, x)
			return nil, err
		}
		a = append(a, b)
	}
	return a, nil
}
func loadSlice(x int, r io.Reader) ([]int, error) {
	a := []int{}
	for i := 0; i < x; i++ {
		x, err := readInt(r)
		if err != nil {
			return nil, err
		}
		a = append(a, x)
	}
	return a, nil
}

func readInt(r io.Reader) (int, error) {
	var b int
    _, err := fmt.Fscan(r, &b)
	if err != nil {
		return 0, err
	}
	return b, nil
}
