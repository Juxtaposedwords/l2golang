package group

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

	return [][]string{}
}

func anagramID(word string) string {
	return word

}
