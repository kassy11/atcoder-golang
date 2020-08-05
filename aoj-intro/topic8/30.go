package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	for{
		fmt.Scan(&input)

		if input == "0"{
			break
		}

		var sum int
		for _, s := range input{
			i, _ := strconv.Atoi(string(s))
			sum += i
		}
		fmt.Println(sum)
	}
}