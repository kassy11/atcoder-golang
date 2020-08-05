package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&nums[i])
	}

	for i:=len(nums)-1;i>=0; i--{
		fmt.Print(nums[i])

		if i == 0{
			fmt.Print("\n")
		}else{
			fmt.Print(" ")
		}
	}
}