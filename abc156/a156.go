package main

import "fmt"

func main() {
	var n, r, rate int
	fmt.Scan(&n, &r)
	if n>=10{
		rate = r
	}else{
		rate = 100 * (10 - n) + r
	}

	fmt.Println(rate)
}