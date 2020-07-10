package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	sum := 0
	for i:=0; i<len(nums); i++{
		fmt.Scan(&nums[i])
		sum += nums[i]
	}

	sort.Ints(nums)
	fmt.Println(nums[0], nums[n-1], sum)
}