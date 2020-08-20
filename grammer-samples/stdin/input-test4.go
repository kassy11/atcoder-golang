package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewScanner(os.Stdin)

func nextInt()int{
	reader.Scan()
	i, e := strconv.Atoi(reader.Text())
	if e != nil{
		panic(e)
	}
	return i
}

func main() {
	reader.Split(bufio.ScanWords)
	n := nextInt()
	x := 0
	for i:=0; i<n; i++{
		x += nextInt()
	}
	fmt.Println(x)

}