package main

import "fmt"

func main() {
	var m string
	fmt.Scan(&m)
	if m[2] == m[3] && m[4] == m[5]{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}