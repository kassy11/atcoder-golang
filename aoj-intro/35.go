package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, tpoint, hpoint int
	fmt.Scan(&n)
	taro := make([]string, n)
	hanako := make([]string, n)
	for i:=0; i<n; i++{
		fmt.Scan(&taro[i], &hanako[i])
		taro[i] = strings.ToLower(taro[i])
		hanako[i] = strings.ToLower(hanako[i])

		if taro[i] > hanako[i]{
			tpoint += 3
		}else if taro[i] == hanako[i]{
			tpoint++
			hpoint++
		}else{
			hpoint += 3
		}
	}

	fmt.Println(tpoint, hpoint)
}