package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, result, distance int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&nums[i])
	}

	sort.Ints(nums)
	for i:=0; i<n; i++{
		result += (nums[i] - nums[0]) * (nums[i] - nums[0])
	}

	for p:=nums[0]; p<=nums[n-1]; p++{
		for i:=0; i<n; i++{
			distance += (nums[i] - p) * (nums[i] - p)
		}
		if distance < result{
			result = distance
		}
		distance = 0
	}

	fmt.Println(result)
}