package anagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sRunes := map[rune]int{}
	for _, rune := range s {
		sRunes[rune]++
	}
	for _, rune := range t {
		if _, ok := sRunes[rune]; !ok {
			return false
		}
		sRunes[rune]--

		if sRunes[rune] == 0 {
			delete(sRunes, rune)
		}
	}
	return len(sRunes) == 0
}
