package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	y := make([]int, n)

	for i:=0; i<len(x); i++{
		fmt.Scan(&x[i])
	}
	for i:=0; i<len(y); i++{
		fmt.Scan(&y[i])
	}

	d1 := 0.0
	d2 := 0.0
	d3 := 0.0
	dinf := 0.0
	for i := 0; i < n; i++ {
		xi := float64(x[i])
		yi := float64(y[i])

		absi := math.Abs(xi - yi)
		d1 += absi
		d2 += math.Pow(absi, 2)
		d3 += math.Pow(absi, 3)
		dinf = math.Max(dinf, absi)
	}

	fmt.Println(d1)
	fmt.Println(math.Sqrt(d2))
	fmt.Println(math.Pow(d3, 1.0/3.0))
	fmt.Println(dinf)
}