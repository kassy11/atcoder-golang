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

	maxv := nums[1] - nums[0]
	minv := nums[0]

	for i:=1; i<len(nums); i++{
		// 最小値との差を考えることで最大値を得る
		maxv = max(maxv, nums[i]- minv)

		// 最小値を記録していく
		minv = min(minv, nums[i])
	}

	fmt.Println(maxv)
}

func max(x, y int)int{
	if x > y{
		return x
	}
	return y
}

func min(x, y int)int{
	if x > y{
		return y
	}
	return x
}