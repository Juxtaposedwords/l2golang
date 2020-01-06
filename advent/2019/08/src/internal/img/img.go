package img

import (
	"bufio"
	//"github.com/google/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
"fmt"
	"io"
	"strconv"
)

type Image struct {
	layers [][][]int
}

// Let's make it so we can read this weird/cryptic text when printing.
func (im Image) pixel(x, y int) string {
	for i := 0; i < len(im.layers); i++ {
		switch im.layers[i][x][y] {
		case 2:
			continue
		case 1:
			return `□`
		case 0:
			return `█`
		}

	}
	return ""
}

func (im Image) print() string {
	if len(im.layers) == 0{
		return ""
	}
	var output string
		output = fmt.Sprintf("%s\n", output)

	for r, row := range im.layers[0] {
		for c := range row {
		output = fmt.Sprintf("%s%s", output, im.pixel(r, c))
		}
		output = fmt.Sprintf("%s\n", output)
	}
	return output
}

// the easiest way to actually solve the problem is to disregard the pixel dimensions.
func layerDigitSums(r io.Reader, columns, rows int) (int, error) {
	/* 
	l, err := layerMatrix(r,columns,rows)
	if err != nil {
		return 0, err
	}
	// herei 
	//img := &Image{layers: l}
	//logger.Infof("%s %#v",img.print(),l)
	*/
	count, err := layerCount(r, columns*rows)
	if err != nil {
		return 0, err
	}

	fewestZeroes := count[0][0]
	for k := range count {
		if _, ok := count[k]; !ok {
			continue
		}
		if count[k][0] < fewestZeroes && k != 0 {
			fewestZeroes = k
		}
	}

	return count[fewestZeroes][1] * count[fewestZeroes][2], nil
}

func layerMatrix(r io.Reader, columns, rows int) ([][][]int, error) {
	digits, err := ints(r)
	if err != nil {
		return nil, err
	}
	area := columns * rows

	if len(digits)%area != 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "invalid dimension for input")
	}
	var layers [][][]int
	for i := 0; i < len(digits); i += area {
		currentLayer := digits[i : i+area]
		var workingLayer [][]int
		for j := 0; j < len(currentLayer); j += columns {
			workingLayer = append(workingLayer, currentLayer[j:j+columns])
		}
		layers = append(layers, workingLayer)
	}
	return layers, nil
}
func ints(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	var output []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		output = append(output, x)
	}
	return output, nil
}
func layerCount(r io.Reader, area int) (map[int]map[int]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	count := map[int]map[int]int{}

	var i int
	for scanner.Scan() {
		layer := int(i / area)
		digit, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		if _, ok := count[layer]; !ok {
			count[layer] = map[int]int{}
		}
		count[layer][digit]++
		i++
	}
	return count, nil
}
