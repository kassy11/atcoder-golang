package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n < m {
		return
	}

	if n == m{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}