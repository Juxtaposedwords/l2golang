package group

import (
	"sort"
)

func GroupAnagrams(strs []string) [][]string {
	anagramMap := map[string][]string{}
	for _, word := range strs {
		id := anagramID(word)

		if _, ok := anagramMap[id]; !ok {
			anagramMap[id] = []string{word}
		} else {
			anagramMap[id] = append(anagramMap[id], word)
		}

	}
	var output [][]string
	for _, v := range anagramMap {
		output = append(output, v)
	}
	return output
}

func anagramID(word string) string {
	if len(word) == 0 {
		return ""
	}
	runeSlice := []rune(word)
	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})
	return string(runeSlice)
}
