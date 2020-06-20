package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := make([]int, m)
	c := make([]int, m)
	for i:=0; i<m; i++{
		fmt.Scan(&s[i], &c[i])
	}

	for i:=0; i<int(math.Pow(10, float64(n))); i++{
		numstr := strconv.Itoa(i)
		for j:=0; j<m; j++{	
			if numstr[nums[j]-1] == c[m]{

			}
		}
	}
}