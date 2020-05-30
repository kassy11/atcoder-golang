package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	var a int
	var s string
	fmt.Scan(&a)
	s = strconv.Itoa(a)
	fmt.Println(strings.Count(s, "1"))
}