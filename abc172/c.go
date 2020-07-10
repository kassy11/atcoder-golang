package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	a := make([]int, n)
	b := make([]int, m)

	for i:=0; i<n; i++{
		fmt.Scan(&a[i])
	}
	for i:=0; i<m; i++{
		fmt.Scan(&b[i])
	}

	sum := 0
	count := 0
	for i:=0;i<len(a);i++{
		for j:=0;j<len(b);j++{
			if sum >= k{
				goto pri
			}
			if a[i] > b[j]{
				sum += b[j]
				b = b[j:]
				count++
			}else{
				sum += a[i]
				a = a[i:]
				count++
			}

		}
	}

	pri:
	fmt.Println(count)
}