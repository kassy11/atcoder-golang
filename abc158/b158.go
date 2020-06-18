package main

import "fmt"

func main() {
	var n, a, b, count int
	fmt.Scan(&n, &a, &b)

	if a == 0{
		count = 0
	}else if a >= n{
		count = n
	}else if a < n{
		if (n % (a+b)) > a{
			count = a * (n / (a+b)) + a
		}else{
			count = a * (n / (a+b)) + (n % (a+b))
		}
	}

	fmt.Println(count)
}