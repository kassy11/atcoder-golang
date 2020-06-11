package main

import (
	"fmt"
)

func main() {
	var a, b, c, ta, tb, tc int
	fmt.Scan(&a, &b, &c)
	origin_a := a
	origin_b := b
	origin_c := c

	count := 0
	flag := true
	for a % 2 == 0 && b % 2 == 0 && c % 2 == 0{
		ta = a
		tb = b
		tc = c
		a = tb / 2 + tc / 2
		b = ta / 2 + tc / 2
		c = ta / 2 + tb / 2

		if a == origin_a && b == origin_b && c == origin_c{
			fmt.Println(-1)
			flag =false
			break
		}

		count++
	}
	if flag{
		fmt.Println(count)
	}
}