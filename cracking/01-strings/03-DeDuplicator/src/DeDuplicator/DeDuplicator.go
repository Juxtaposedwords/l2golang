package DeDuplicatorString

func DeDuplicator(s string) string {
	c := make(map[rune]int)
	i := 0
	for _, e := range []rune(s) {
		if _, ok := c[e]; ok {
			continue
		}
		c[e] = i
		i++
	}
	x := make([]rune, len(c), len(c))
	for k, v := range c {
		x[v] = k
	}
	return string(x)
}
