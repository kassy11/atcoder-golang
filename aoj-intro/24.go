package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	a := make([][]int, n)
	b := make([][]int, m)
	c := make([][]int, n)

	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	for j := 0; j < m; j++ {
		b[j] = make([]int, 1)
		for k := 0; k < 1; k++ {
			fmt.Scanf("%d", &b[j][k])
		}
	}
	for i := 0; i < n; i++ {
		c[i] = make([]int, 1)
		for k := 0; k < 1; k++ {
			for j := 0; j < m; j++ {
				c[i][k] += a[i][j] * b[j][k]
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(c[i][0])
	}
}
