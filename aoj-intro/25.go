package main

import "fmt"

func main() {
	for {
		var m, n, r int
		fmt.Scanf("%d %d %d", &m, &n, &r)
		if m == -1 && n == -1 && r == -1 {
			break
		}
		if F(m,n,r) {
		}else if A(m, n, r) {
		} else if B(m, n, r) {
		} else if C(m, n, r) {
		} else if D(m, n, r) {
		}
	}
}

func A(m, n, r int) bool {
	if m+n >= 80 {
		fmt.Println("A")
		return true
	}
	return false
}

func B(m, n, r int) bool {
	if m+n < 80 && m+n >= 65 {
		fmt.Println("B")
		return true
	}
	return false
}

func C(m, n, r int) bool {
	if m+n < 65 && m+n >= 50 {
		fmt.Println("C")
		return true
	} else if m+n < 50 && m+n >= 30 && r >= 50 {
		fmt.Println("C")
		return true
	}
	return false
}

func D(m, n, r int) bool {
	if m+n < 50 && m+n >= 30 {
		fmt.Println("D")
		return true
	}
	return false
}

func F(m, n, r int) bool {
	if m == -1 || n == -1 || m+n < 30 {
		fmt.Println("F")
		return true
	}
	return false
}
