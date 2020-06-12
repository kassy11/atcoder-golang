package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	child := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&child[i])
	}
	sort.Ints(child)

	count := 0
	for l:=0; l<n; l++{
		if x >= child[l] && x >= 0 {
			x -= child[l]
			count++
		}
	}
	fmt.Println(count)
}