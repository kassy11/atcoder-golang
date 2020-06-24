package main

import "fmt"

func main() {
	var n, m int
	p := make([]int, m)
	s := make([]string, m)
	for i:=0; i<m; i++{
		fmt.Scan(&p[i])
		if p[i] > n || p[i] < 0{
			return
		}
		fmt.Scan(&s[i])
		if s[i] != "AC" || s[i] != "WA"{
			return
		}
	}

	

}