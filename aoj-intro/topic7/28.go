package main

import "fmt"

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)

	a := make([][]int, n)
	b := make([][]int, m)
	c := make([][]int, n)

	for i:=0; i<len(a); i++{
		a[i] = make([]int, m)
		for j:=0; j<m; j++{
			fmt.Scan(&a[i][j])
		}
	}
	for i:=0; i<len(b); i++{
		b[i] = make([]int, l)
		for j:=0; j<l; j++{
			fmt.Scan(&b[i][j])
		}
	}
	for i := 0; i < n; i++ {
		c[i] = make([]int, l)
		for k := 0; k < l; k++ {
			for j := 0; j < m; j++ {
				c[i][k] += a[i][j] * b[j][k]
			}
		}
	}
	for i := 0; i < n; i++ {
		for k:=0;k < l;k++ {
			fmt.Print(c[i][k])
			if k != l-1{
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}