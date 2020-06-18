package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result := -1
	for i:=1; i<=1009; i++{
		if int(float64(i) * 0.08) == a && int(float64(i) * 0.1) == b{
			result = i
			break
		}
	}

	fmt.Println(result)
}