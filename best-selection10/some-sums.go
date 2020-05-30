package main

import "fmt"

func sumOfDegit(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main(){
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	result := 0
	for i:=1; i<n+1; i++{
		digit_num := sumOfDegit(i)
		if digit_num >= a && digit_num <= b{
			result += i
		}
	}
	fmt.Println(result)
}