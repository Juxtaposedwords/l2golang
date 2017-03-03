package ReverseString

func Reverse(s string) string {
	m := len(s) / 2
	j := len(s) - 1
	x := []rune(s)
	for i := 0; i < m; i++ {
		x[i], x[j-i] = x[j-i], x[i]
	}
	return string(x)
}
