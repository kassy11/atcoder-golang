package main

import "fmt"

func main(){
	var n, m int
	fmt.Scan(&n, &m)
	matrix := make([][]int, n)
	for i:=0; i<len(matrix); i++{
		matrix[i] = make([]int, m)
		for j:=0; j<len(matrix[i]); j++{
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Println(matrix)
}