package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var A, B, X int
	fmt.Scan(&A, &B, &X)
	N := sort.Search(1e9+1, func(i int) bool {
		if i == 0 {
			return false
		}
		return A*i+B*len(strconv.Itoa(i)) > X
	})
	fmt.Println(N - 1)
}
