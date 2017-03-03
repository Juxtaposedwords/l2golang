package uniqueString

func Unique(s string) bool {
	x := make(map[rune]bool)
	for _, e := range s {
		if _, ok := x[e]; ok {
			return false
		}
		x[e] = false
	}
	return true
}
