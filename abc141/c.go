package main

import "fmt"

func main() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)
	seq := make([]int, q)
	point := make([]int, n)
	for i:=0; i<len(point); i++{
		point[i] = k
	}

	for i:=0; i<len(seq); i++{
		fmt.Scan(&seq[i])
		for j:=0; j<len(point); j++{
			if j != seq[i]-1{
				point[j]--
			}
		}
	}

	fmt.Println("-----")
	for i:=0; i<len(point); i++{
		if point[i] >= 1{
			fmt.Println("Yes")
		}else {
			fmt.Println("No")
		}
	}


}