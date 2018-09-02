package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	x := []string{
		"596d",
		"6c30",
		"4c6d",
		"7835",
		"4c7a",
		"4a32",
		"6258",
		"5a45",
		"5a6a",
		"514b"}
	var output string
	for _, char := range x {
		c, err := hex.DecodeString(char)
		if err != nil {
			fmt.Printf("Error deocding %s: %s", char, err)
			continueq
		}
		output += fmt.Sprintf("%s", c)

	}
	o, err := base64.StdEncoding.DecodeString(output)
	output = string(o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
