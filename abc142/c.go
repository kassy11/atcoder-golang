package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	ans := make([]int, n)
	for i, _ := range nums{
		fmt.Scan(&nums[i])
	}

	for i:=1; i<=n; i++{
		for j:=0; j<n; j++{
			if i == nums[j]{
				ans[i-1] = j+1
				break
			}
		}
	}

	for _, v := range ans{
		fmt.Println(v)
	}
}