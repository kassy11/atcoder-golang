package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for l:=0; l<n-1; l++{
		fmt.Scan(&nums[l])
	}

	count_map := make(map[int]int)

	for i:=1; i<=n; i++{
		for m:=0; m<n-1; m++{
			if i == nums[m]{
				count_map[i] ++
			}
		}
		fmt.Println(count_map[i])
	}

}