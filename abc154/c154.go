package main

import "fmt"

func main(){

	matrix := make([][]string, n)
	for i:=0; i<len(matrix); i++{
		matrix[i] = make([]string, m)
		for j:=0; j<len(matrix[i]); j++{
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Println(matrix)
}