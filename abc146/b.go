package main

import "fmt"

func main() {
	var str string
	var n int32
	fmt.Scan(&n)
	fmt.Scan(&str)

	for _, s := range str{
		x := s - 'A'
		x = (x+n)%26
		fmt.Print(string('A'+x))
	}
	fmt.Println("")
}