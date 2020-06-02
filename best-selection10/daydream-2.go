package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	for len(s) > 0 {
		switch {
		case strings.HasSuffix(s, "dreamer"):
			s = s[:len(s)-7]
		case strings.HasSuffix(s, "dream"):
			s = s[:len(s)-5]
		case strings.HasSuffix(s, "eraser"):
			s = s[:len(s)-6]
		case strings.HasSuffix(s, "erase"):
			s = s[:len(s)-5]
		default:
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}