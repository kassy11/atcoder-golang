package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	for i:=0; i<=n; i++{
		for j:=0; j<=n-i; j++{
			for k:=0; k<=n-i-j; k++{
				if 10000*i+5000*j+1000*k == y{
					fmt.Println(i, j, k)
				}
			}
		}
	}
	fmt.Println(-1, -1, -1)
}