package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0; i<len(nums); i++{
		fmt.Scan(&nums[i])
	}

	count := 0
	sort.Ints(nums)

	for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++{
			for k:=j+1; k<len(nums); k++{
				if nums[i] != nums[j] && nums[j] != nums[k] && nums[i] != nums[k]{
					if nums[i] + nums[j] > nums[k]{
						count++
					}
				}
			}
		}
	}

	fmt.Println(count)
}