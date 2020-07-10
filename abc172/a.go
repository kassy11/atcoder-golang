package main

import "fmt"

func main() {
	var a int
	if a < 0 || a > 10 {
		return
	}
	fmt.Scan(&a)
	fmt.Println(a + a * a + a * a * a)
}