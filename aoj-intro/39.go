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
		sum := 0
		for i:=0; i<len(s); i++{
			fmt.Scan(&s[i])
			sum += s[i]
		}
		average := sum / n

		bsum := 0
		for i:=0; i<len(s); i++{
			bsum += (s[i]-average) * (s[i]-average)
		}
		fmt.Println(math.Sqrt(float64(bsum / n)))
	}
}