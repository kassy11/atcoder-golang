// ここからトピック10
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(math.Sqrt(math.Pow((a-c), 2) + math.Pow((b-d), 2)))
}