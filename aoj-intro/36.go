package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var str string
	var n int
	fmt.Scan(&str)
	fmt.Scan(&n)

	for i:=0; i<n; i++{
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		cmd := strings.Split(scanner.Text(), " ")
		switch  cmd[0]{
		case "replace":
			str =
		case "reverse":

		case "print":
			fmt.Println(str[strconv.Atoi(cmd[0])])
		}

	}
}