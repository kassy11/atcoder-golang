package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0

	count += n / 25
	n = n % 25
	count += n/10
	n = n % 10
	count += n / 5
	n = n % 5
	count += n

	fmt.Println(count)
}