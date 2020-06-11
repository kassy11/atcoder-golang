package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	if h == 1 || w == 1{
		fmt.Println(1)
	}else{
		fmt.Println((h * w + 1) / 2 )
	}


}