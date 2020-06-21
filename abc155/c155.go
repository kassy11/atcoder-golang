package main

import (
	"fmt"
)


func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	result := make(map[string]int)
	for i:=0; i<len(s); i++{
		fmt.Scan(&s[i])
		result[s[i]]++
	}
	// mapで文字の出現回数を管理するところまで


}