package main

import "fmt"

func main() {
	var n, k, m, result int
	fmt.Scan(&n, &k, &m)
	a := make([]int, n-1)
	sum := 0
	for i:=0; i<n-1; i++{
		fmt.Scan(&a[i])
		if a[i] > k || a[i] < 0 {
			return
		}
		sum += a[i]
	}

	if sum / n >= m{
		result = 0
	}else if (sum + k) / n < m{
		result = -1
	}else{
		result = (m * n) - sum
	}

	fmt.Println(result)
}