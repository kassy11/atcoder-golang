package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	fmt.Printf("%f %f\n", n * n * math.Pi, 2 * n * math.Pi)
}