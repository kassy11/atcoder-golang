package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&nums[i])
	}

	sort.Ints(nums)

	for l:=nums[0]; l<=nums[n]; l++{
		
	}
}