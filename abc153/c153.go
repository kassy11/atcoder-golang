package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	h := make([]int, n)
	for i := 0; i < len(h); i++ {
		fmt.Scan(&h[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(h)))

	sum := 0
	for i:=k; i<len(h); i++{
		sum += h[i]
	}

	fmt.Println(sum)
}
