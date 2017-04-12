package beautifulWord

import (
	"fmt"
)

const (
	beautTrue  = "Yes"
	beautFalse = "No"
)

var (
	vowels = []rune{'a', 'e', 'i', 'o', 'u', 'y'}
)

func IsBeautiful(input string) string {
	var lastRune rune
	for index, character := range input {
		if index == 0 {
			continue
		}
		lastRune = []rune(input)[index-1]
		if character == lastRune {
			return beautFalse
		}

		if index > 0 && isVowel([]rune(input)[index]) && isVowel([]rune(input)[index-1]) {
			fmt.Printf("%s\n", string(character))
			return beautFalse
		}
	}
	return beautTrue
}

func isVowel(char rune) bool {
	for _, vowel := range vowels {
		if vowel == char {
			return true
		}
	}
	return false
}
