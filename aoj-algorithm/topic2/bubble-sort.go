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
	flag := true
	for flag{
		flag = false
		for i := len(nums)-1; i>0; i--{
			if nums[i] < nums[i-1]{
				tmp := nums[i-1]
				nums[i-1] = nums[i]
				nums[i] = tmp
				flag = true
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

