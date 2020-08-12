package main

import "fmt"

func insertionsort(nums []int){
	for i:=1; i<len(nums); i++{
		v := nums[i]
		j := i-1

		for j>=0 && nums[j]>v{
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = v
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0; i<len(nums); i++{
		fmt.Scan(&nums[i])
	}

	insertionsort(nums)
}