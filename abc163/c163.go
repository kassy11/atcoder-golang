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

	for _, v := range nums{
		count_map[v] ++
	}

	for i := 1; i <= n; i++ {
		fmt.Println(count_map[i])
	}
}