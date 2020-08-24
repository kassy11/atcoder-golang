package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 0{
		fmt.Println(float64(n/2) / float64(n))
	}else{
		fmt.Println(float64(n/2+1)/float64(n))
	}
}