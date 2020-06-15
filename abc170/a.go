package main

import "fmt"

func main() {
	var result int
	nums := make([]int, 5)
	for i:=0; i<5; i++{
		fmt.Scan(&nums[i])

		if nums[i] == 0{
			result = i+1
		}
	}

	fmt.Println(result)

}