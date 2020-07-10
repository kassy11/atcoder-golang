package main

import "fmt"

func main() {
	n := make([]int, 0)
	for i:=0; ;i++{
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		n = append(n, num)
	}

	for i:=0; i<len(n); i++{
		fmt.Printf("Case %d: %d\n", i+1, n[i])
	}
}