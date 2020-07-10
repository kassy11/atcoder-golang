package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := make([]int, 3)
	for i:=0; i<3; i++{
		fmt.Scan(&nums[i])
	}

	sort.Ints(nums)
	fmt.Printf("%d ", nums[0])
	fmt.Printf("%d ", nums[1])
	fmt.Printf("%d\n", nums[2])

}