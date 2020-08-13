package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(gcd(x, y))
}

func gcd(a, b int)int{
	if a<b{
		tmp := a
		a = b
		b = tmp
	}

	r := a % b
	for r != 0{
		a = b
		b = r
		r = a % b
	}
	return b
}