package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)

	for i:=0; i<len(nums); i++{
		fmt.Scan(&nums[i])
	}

	max := nums[1] - nums[0]
	for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++{
			if nums[j]-nums[i] > max{
				max = nums[j] - nums[i]
			}
		}
	}
	fmt.Println(max)
}