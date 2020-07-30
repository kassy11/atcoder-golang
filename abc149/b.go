package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)
	if k<=a{
		a = a - k
	}else if k<=a+b{
		b = b - (k-a)
		a = 0
	}else{
		a = 0
		b = 0
	}
	fmt.Println(a, b)
}