package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n-1)

	for i:=0; i<len(a); i++{
		fmt.Scan(&a[i])
	}
	for i:=0; i<len(b); i++{
		fmt.Scan(&b[i])
	}
	for i:=0; i<len(c); i++{
		fmt.Scan(&c[i])
	}

	point := b[a[0]-1]
	for i:=1; i<len(a); i++{
		point += b[a[i]-1]
		if a[i-1] + 1 == a[i]{
			point += c[a[i-1]-1]
		}
	}
	fmt.Println(point)
}