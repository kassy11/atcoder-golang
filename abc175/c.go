package main

import (
	"fmt"
	"math"
)

func main() {
	var x, k, d int
	fmt.Scan(&x, &k, &d)
	var min int

	min = 1000000000
	for i:=0; i<k; i++{
		if math.Abs(float64(x) - float64(d)) < float64(min){
			min = int(math.Abs(float64(x) - float64(d)))
			x = x -d
		}else if math.Abs(float64(x) + float64(d)) < float64(min){
			min = int(math.Abs(float64(x) + float64(d)))
			x = x + d
		}
	}
	fmt.Println(min)
}