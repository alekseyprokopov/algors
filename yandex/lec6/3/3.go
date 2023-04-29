package main

import "fmt"

func main() {
	var w, h, n int
	fmt.Scanf("%d %d %d", &w, &h, &n)
	result := rBinSearch(w, h, n)
	fmt.Println(result)
}

func rBinSearch(w, h, n int) int {
	l := 1
	r := min(h, w) * n

	for l < r {
		m := (r + l) / 2
		if (m/w)*(m/h) >= n {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
