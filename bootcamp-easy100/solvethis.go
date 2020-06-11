package main

import "fmt"

func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)
	b := make([]int, m)
	for i:=0; i<m; i++{
		fmt.Scan(&b[i])
	}

	// 二次元配列aの初期化
	a := make([][]int, n)
	for l:=0; l<len(a); l++{
		a[l] = make([]int, m)
		for j:=0; j<len(a[l]); j++{
			fmt.Scan(&a[l][j])
		}
	}

	result := make([]int, n)
	count := 0
	for q:=0; q<len(a); q++{
		for s:=0; s<len(a[q]); s++{
			result[q] += a[q][s] * b[s]
		}
		if result[q]+c > 0{
			count++
		}
	}
	fmt.Println(count)
}