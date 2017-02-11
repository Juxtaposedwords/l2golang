package main 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func readInput(t string) (*bufio.Scanner, error){
	st, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if (st.Mode() & os.ModeCharDevice) == 0 {
		return bufio.NewScanner(os.Stdin), nil
	} else {
		f, err := os.Open(t)
		if err != nil  {
			return nil, err
		}
		return bufio.NewScanner(f), nil
	}
}

func main(){
	scanner, err := readInput("GreenHills.txt")
	if err != nil {
		fmt.Println(err)
	}
	for scanner.Scan(){
		fmt.Println(strings.Replace(scanner.Text(), "Earth", "Mars", -1))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}