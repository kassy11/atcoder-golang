package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	b := make([]int, n-1)
	for i:=0; i<len(b); i++{
		fmt.Scan(&b[i])
	}
	
}