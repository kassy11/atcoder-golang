package main

import "fmt"

func main(){
	var a int
	var b float64
	fmt.Scan(&a, &b)
	fmt.Println(int64(float64(a)*b))
}