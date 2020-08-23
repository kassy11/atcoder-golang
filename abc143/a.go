package main

import "fmt"

func main() {
	var a, b, ans int
	fmt.Scan(&a, &b)

	if b * 2 >= a{
		ans = 0
	}else{
		ans = a - b * 2
	}

	fmt.Println(ans)
}