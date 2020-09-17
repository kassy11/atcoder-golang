package main

import (
	"fmt"
)

const mod = 1000000007

func powmod(x int, y int)int{
	res := 1
	for i:=0; i<y; i++{
		res = res*x%mod
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	ans := powmod(10, n) - powmod(9, n) - powmod(9, mod) + powmod(8, mod)
	ans%=mod
	ans=(ans+mod)%mod
	fmt.Println(ans)
}