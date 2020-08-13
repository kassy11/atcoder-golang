package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	min := n
	for i:=1; i<=n; i++{
		for j:=1; j<=n; j++{
			if i*j == n && (i-1)+(j-1) < min{
				min = (i-1)+(j-1)
				break
			}
		}
	}

	fmt.Println(min)
}