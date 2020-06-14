package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if len(s) < 3 || len(s) > 20{
		return
	}
	fmt.Println(s[0:3])
}