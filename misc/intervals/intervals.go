package main

type tuple struct {
	start int
	end int
}

func theFunc(b []tuple, n tuple) []tuple{
	var r []tuple
	first, last := b[0], b[len(b)-1]
	if n.end < first.start {
		r = append([]tuple{n},b...)
		return r
	}
	if n.start > last.end {
	}
	return r
}