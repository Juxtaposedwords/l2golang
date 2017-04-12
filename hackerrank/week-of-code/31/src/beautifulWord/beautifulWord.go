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
		// All beautiful checks are based off the relationships of characters to previous characters, so skip the first one
		if index == 0 {
			continue
		}
		lastRune = []rune(input)[index-1]
		// check to see if the previous rune and present rune are the same
		if character == lastRune {
			return beautFalse
		}
		// check to see if the previous rune and the present rune are both variables
		if isVowel(character) && isVowel(lastRune) {
			fmt.Printf("%s\n", string(character))
			return beautFalse
		}
	}
	// If none of the checks have been provided, then return True
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
