package main

import (
	"fmt"
	"math"
	"strconv"
)

func isInteger(x float64) bool {
	return math.Floor(x) == x
}
func main() {
	var a, b string
	fmt.Scan(&a, &b)
	num, _ := strconv.Atoi(a+b)
	if isInteger(math.Sqrt(float64(num))){
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}