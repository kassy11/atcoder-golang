package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	var high_point = num / 500
	var low_poinnt = (num%500) / 5
	fmt.Println(1000 * high_point + 5 * low_poinnt)
}