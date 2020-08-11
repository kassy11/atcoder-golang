package main

import (
	"fmt"
	"math"
)

func main() {
	for{
		var n int
		fmt.Scan(&n)
		if n == 0{
			break
		}

		s := make([]int, n)
		sum := 0.0
		for i:=0; i<len(s); i++{
			fmt.Scan(&s[i])
			sum += float64(s[i])
		}
		average := sum / float64(n)

		bsum := 0.0
		for i:=0; i<len(s); i++{
			bsum += (float64(s[i])-average) * (float64(s[i])-average)
		}
		fmt.Println(math.Sqrt(bsum / float64(n)))
	}
}