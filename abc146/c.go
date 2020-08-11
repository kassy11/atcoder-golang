package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	max := 0
	buy := 0
	for i:=1; i<1000000000; i++{
		len := len(strconv.Itoa(i))
		num := a*i + b*len
		//fmt.Println(len, i)
		//fmt.Println(a*i + b*len)
		if num >x{
			break
		}
		if num > max {
			max = num
			buy = i
		}
	}
	fmt.Println(buy)
}