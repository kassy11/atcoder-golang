package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	nums := make([]int, n)

	count := 0
	for i, _ := range nums{
		fmt.Scan(&nums[i])
		if nums[i] >= k{
			count++
		}
	}

	fmt.Println(count)
}