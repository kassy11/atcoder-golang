package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)

	count := 0
	for i:=0; i<len(nums); i++{
		fmt.Scan(&nums[i])
		if primenum(nums[i]){
			count++
		}
	}
	fmt.Println(count)
}

func primenum(x int)bool{
	for i := 2; i <= int(math.Sqrt(float64(x))); i++{
		if x % i == 0{
			return false
		}
	}

	return true
}