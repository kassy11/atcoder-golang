
package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextLine() string{
	scanner.Scan()
	return scanner.Text()
}

func main() {
	n := nextLine()
	num, _ := strconv.Atoi(n)
	arry := make([]int, num)

	for i:=0; i<len(arry); i++{
		arry[i], _ = strconv.Atoi(nextLine())
	}

	fmt.Println(arry)
}