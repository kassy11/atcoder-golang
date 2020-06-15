package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	ans := "No"
	for i:=0; i<x+1; i++{
		l := x - i
		if 2 * i + 4 * l == y{
			ans = "Yes"
		}

	}

	fmt.Println(ans)
}