package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mi := int(1e15 + 1)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			mi = min(mi, i+n/i)
		}
	}
	fmt.Println(mi - 2)
}

func min(a, b int)int{
	if a > b{
		return b
	}
	return a
}