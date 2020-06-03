package main

import "fmt"

func swap(m, n int) (int, int) {
	var i, j = n, m
	return i, j
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	a, b = swap(a, b)
	a, c = swap(a, c)
	fmt.Println(a, b, c)
}