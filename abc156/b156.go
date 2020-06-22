package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for n > 0{
		n = n / 10
		count++
	}

	fmt.Println(count)
}