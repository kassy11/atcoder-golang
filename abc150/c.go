package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	q := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&p[i])
		fmt.Scan(&q[i])
	}

	
}