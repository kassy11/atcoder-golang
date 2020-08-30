package main

import "fmt"

func main() {
	 var d, t, s float64
	 fmt.Scan(&d, &t, &s)

	 if t >= d/s{
	 	fmt.Println("Yes")
	 }else{
		fmt.Println("No")
	 }
}