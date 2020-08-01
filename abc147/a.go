package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a + b + c >= 22{
		fmt.Println("bust")
	}else{
		fmt.Println("win")
	}
}