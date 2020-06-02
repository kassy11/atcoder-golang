package main

import (
	"fmt"
)

func sliceUnique(target []int)(unique []int){
	m := map[int]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func main(){
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&nums[i])
	}

	unique := sliceUnique(nums)
	fmt.Println(len(unique))
}