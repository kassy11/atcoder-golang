package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0; i<len(nums); i++{
		fmt.Scan(&nums[i])
	}
	bubbleSort(nums)
}

func bubbleSort(nums []int){
	count := 0
	for i:=0; i<len(nums); i++{
		for j:=len(nums)-1; j>i; j--{
			if nums[j-1] > nums[j]{
				tmp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = tmp
				count++
			}
		}
	}

	for k:=0; k<len(nums); k++{
		fmt.Print(nums[k])
		if k != len(nums)-1{
			fmt.Print(" ")
		}
	}
	fmt.Println("")
	fmt.Println(count)
}

