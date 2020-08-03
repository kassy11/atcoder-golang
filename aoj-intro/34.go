package main

import "fmt"

func main() {
	for{
		var str string
		var m int
		fmt.Scan(&str)
		if str == "-"{
			break
		}
		fmt.Scan(&m)
		h := make([]int, m)
		for i:=0; i<len(h); i++{
			fmt.Scan(&h[i])
			str = str[h[i]:] + str[:h[i]]
		}
		fmt.Println(str)
	}
}