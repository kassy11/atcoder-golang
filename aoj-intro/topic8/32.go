package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, p string
	fmt.Scan(&s)
	fmt.Scan(&p)
	if strings.Contains(strings.Repeat(s, 2), p){
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}