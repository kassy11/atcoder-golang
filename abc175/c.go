package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x, k, d int
	fmt.Scan(&x, &k, &d)

	fmt.Println(solve(x, k, d))
}

func solve(X, K, D int) int {
	if X < 0 {
		X = -X
	}

	if cmp(X, K, D) {
		return X - K*D
	}

	k := X / D
	x := X - D*k
	d := K - k
	if d%2 == 0 {
		return x
	} else {
		return D - x
	}
}

func cmp(X, K, D int) bool {
	x := big.NewInt(int64(X))
	k := big.NewInt(int64(K))
	d := big.NewInt(int64(D))

	return 0 < x.Cmp(k.Mul(k, d))
}
