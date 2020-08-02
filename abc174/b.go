package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	count := 0
	x := make([]int, n)
	y := make([]int, n)

	for i:=0; i<len(x); i++{
		fmt.Scan(&x[i])
		fmt.Scan(&y[i])
		if x[i] > d || y[i] > d{
			continue
		}

		if x[i] * x[i] + y[i] * y[i] <= d * d{
			count++
		}

	}

	fmt.Println(count)
}