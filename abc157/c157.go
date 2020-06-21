package main

import (
	"fmt"
	"strconv"
)

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}


func intToString(i []int) []string {
	f := make([]string, len(i))
	for n := range i {
		f[n] = strconv.Itoa(i[n])
	}
	return f
}

func main() {
	var n, m int
	var numstr string
	result := -1
	fmt.Scan(&n, &m)

	s := make([]int, m)
	c := make([]int, m)
	array := make([]int, n)
	for i:=0; i<n; i++{
		array[i] = 9
	}
	for i:=0; i<m; i++{
		fmt.Scan(&s[i], &c[i])
	}


	if n<=m && contains(s, 1){
		// n桁の数字について
		for i:=1; i<=n; i++ {
			if !contains(s, i){
				array[i-1] = 0
			}
			// s[l]と値が等しく、そのcの値が最小なら追加
			for l := 0; l < m; l++ {
				if i == s[l] && c[l] < array[i-1] {
					array[i-1] = c[l]
				}
			}
		}
		stringarry := intToString(array)
		for _, v := range stringarry{
			numstr += v
		}
		result, _ = strconv.Atoi(numstr)
	}

	fmt.Println(result)
}