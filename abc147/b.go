package main

import "fmt"

func main() {
	count := 0
	var s string
	fmt.Scan(&s)

	for i,j:=0, len(s)-1; i<j; i, j = i+1, j-1{
		if s[i] != s[j]{
			count++
		}
	}

	fmt.Println(count)
}