package main

import (
	"fmt"
)

func chocoday(x int, d int)int{
	chocoday := 0
	for a:=0; a*x+1<=d; a++{
		chocoday++
	}
	return chocoday
}

func main() {
	var n, d, x int
	fmt.Scan(&n)
	fmt.Scan(&d, &x)
	chococap := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&chococap[i])
	}
	choconum := 0
	for l:=0; l<n; l++{
		choconum += chocoday(chococap[l], d)
	}
	fmt.Println(choconum + x)
}