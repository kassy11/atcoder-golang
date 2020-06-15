package main

import (
	"fmt"
)

func combinationTwo(a int)int{
	return a * (a-1) / 2
}

func main() {
	var even, odd int
	fmt.Scan(&even, &odd)
	// even C 2 + odd C 2の組み合わせが答え

	ans := 0
	if even >= 2 && odd >= 2{
		ans = combinationTwo(even) + combinationTwo(odd)
	}else if even >= 2 && odd < 2{
		ans = combinationTwo(even)
	}else if even < 2 && odd >= 2{
		ans = combinationTwo(odd)
	}

	fmt.Println(ans)
}