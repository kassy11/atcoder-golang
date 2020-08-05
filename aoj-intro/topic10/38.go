package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	s := 0.5 * a * b * math.Sin((c*math.Pi)/float64(180))
	l := a + b + math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)-2*a*b*math.Cos((c*math.Pi)/float64(180)))
	h := b * math.Sin((c*math.Pi)/float64(180))
	fmt.Println(s, l, h)
}