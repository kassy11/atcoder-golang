package main

import (
	"fmt"
	"sort"
)

func main(){
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&nums[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	count := 1
	for i:=1; i<n; i++{
		if nums[i] == nums[i-1]{
			continue
		}
		count++
	}

	fmt.Println(count)

}