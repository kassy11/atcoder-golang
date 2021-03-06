package main

import "fmt"

func checkDigit(num int) bool{
	for num >= 0{
		if num % 10 == 7{
			return true
		}
		num = num / 10
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)

	if checkDigit(n) {
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}

}