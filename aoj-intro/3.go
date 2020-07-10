package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	area := a * b
	role := 2 * a + 2 * b
	fmt.Println(area, role)
}