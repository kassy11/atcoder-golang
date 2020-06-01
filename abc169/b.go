package main

import (
	"fmt"
	"math"
)

func main(){
	var n int
	var result uint64
	result = 1
	fmt.Scan(&n)
	if n < 2{
		fmt.Println(-1)
	}
	nums := make([]int, n)

	for i:=0; i<n; i++{
		fmt.Scan(&nums[i])
	}

	for k:=0; k<n; k++{
		result *= uint64(nums[k])
	}

	if result > uint64(math.Pow10(18)){
		fmt.Println(-1)
	}else{
		fmt.Println(result)
	}

}