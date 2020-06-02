package main

import (
	"fmt"
	"sort"
)

func main(){
	var (
		n int
		alice = 0
		bob = 0
	)
	fmt.Scan(&n)

	cards := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&cards[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cards)))

	for m:=0; m<n; m=m+2{
		alice += cards[m]
		if m+1<n{
			bob += cards[m+1]
		}
	}

	fmt.Println(alice-bob)
}