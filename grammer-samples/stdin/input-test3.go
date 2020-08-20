package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewScanner(os.Stdin)

func nextLine()string{
	reader.Scan()
	return reader.Text()
}

func main() {
	var s string

	s = nextLine()

	fmt.Println(s)
}