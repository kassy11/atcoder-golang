package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	goods := make([]int, n)

	count := 1
	for i:=0; i<n; i++{
		fmt.Scan(&goods[i])
	}

	sort.Ints(goods)

	for l:=1; l<n; l++{
		if goods[l-1] != goods[l] {
			count++
		}
	}

	fmt.Println(count)

}