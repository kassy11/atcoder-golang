package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if k > n && (k- n) > n {
		fmt.Println(n)
	}else{
		t1 := n % k
		t2 := k - t1
		if t1 > t2 {
			fmt.Println(t2)
		}else{
			fmt.Println(t1)
		}
	}
}