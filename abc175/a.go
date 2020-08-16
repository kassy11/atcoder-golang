package main

import (
	"fmt"
)

func main() {
	var str string
	var count int
	fmt.Scan(&str)

	if len(str) != 3 {
		return
	}

	if str == "RRR"{
		count = 3
	}else if str == "SSS" {
		count = 0
	}else if str == "RRS" || str == "SRR"{
		count = 2
	}else{
		count = 1
	}
	fmt.Println(count)
}