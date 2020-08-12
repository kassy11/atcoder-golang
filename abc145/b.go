package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	if len(s) != n{
		return
	}

	ans := "No"
	if n == 1 || n % 2 == 0{
		if s[:n/2] == s[n/2:]{
			ans = "Yes"
		}
	}

	fmt.Println(ans)
}