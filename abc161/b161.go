package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	points := make([]int, n)

	for i:=0; i<n; i++{
		fmt.Scan(&points[i])
	}
	sum := 0
	for _, point := range points{
		sum += point
	}
	chosenable := sum / 4 * m

	sort.Sort(sort.Reverse(sort.IntSlice(points)))

	count := 0
	for l:=0; l<m; l++{
		if points[l] > chosenable{
			count ++
		}
	}

	if count == m{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}


}