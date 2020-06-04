package main

import (
	"fmt"
	"sort"
)

func sliceUnique(target []int) (unique []int) {
	m := map[int]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if m > n {
		return
	}
	points := make([]int, n)
	sum := 0

	for i:=0; i<n; i++{
		fmt.Scan(&points[i])
		sum += points[i]
	}


	sort.Sort(sort.Reverse(sort.IntSlice(points)))
	unique_points := sliceUnique(points)

	count := 0
	for i := 0; i < n; i++ {
		if unique_points[i]*4*m >= sum {
			count++
		}
	}
	if count >= m {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")


}