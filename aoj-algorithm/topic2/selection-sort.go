package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0; i<len(nums); i++{
		fmt.Scan(&nums[i])
	}

	selectionSort(nums)
}

func selectionSort(nums []int){
	count := 0

	for i := 0; i < len(nums); i++ {
		minj := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minj] {
				minj = j
			}
		}
		if i != minj {
			nums[i], nums[minj] = nums[minj], nums[i]
			count++
		}
	}

	for i:=0; i<len(nums); i++{
		fmt.Print(nums[i])
		if i == len(nums) -1{
			fmt.Println("")
		}else{
			fmt.Print(" ")
		}
	}
	fmt.Println(count)
}
