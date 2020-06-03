package main

import (
	"fmt"
	"sort"
)

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	nums := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)

	maxnum := 0
	for l:=0; l<n-1; l++{
		if (nums[l+1] - nums[l]) >= maxnum{
			maxnum = nums[l+1] - nums[l]
		}
	}
	if (k-nums[n-1]) + nums[0] >= maxnum{
		maxnum = (k-nums[n-1]) + nums[0]
	}
	fmt.Println(k - maxnum)

}