package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i:=1; i<=n; i++{
		if (i % 3) != 0 && (i % 5) != 0{
			count += i
		}
	}
	fmt.Println(count)
}