package main

import (
	"fmt"
)

// bitmap we'll use to map out the presence of numbers
type Bitmap []byte

func NewBitmap(max int, filePath string) Bitmap {
	b := make(Bitmap, max, max)
	return b
}

// sets the integer number's present in the bitmap
func (bits Bitmap) Set(i uint) {
	bits[i/8] &= 1 << i%8
	fmt.Println(bits[i/8])
	//return bits[i/8] & i%8 != 0
}

// Not sure here?
func (bits *Bitmap) Get(number int) bool {
	return true
}


func main() {
	test := NewBitmap(13, "bad.txt")
	test.Set(12)
	fmt.Println(test)
}
