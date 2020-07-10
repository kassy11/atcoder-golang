package main

import "fmt"

func main() {
	a := make([]int, 1)
	b := make([]int, 1)
	op := make([]string, 1)

	for {
		var anum, bnum int
		var opsample string
		fmt.Scan(&anum)
		fmt.Scan(&opsample)
		fmt.Scan(&bnum)

		if opsample == "?"{
			break
		}

		a = append(a, anum)
		b = append(b, bnum)
		op = append(op, opsample)
	}

	for i:=0; i<len(a); i++{
		switch op[i] {
		case "+":
			fmt.Println(a[i] + b[i])
		case "-":
			fmt.Println(a[i] - b[i])
		case "*":
			fmt.Println(a[i] * b[i])
		case "/":
			fmt.Println(a[i] / b[i])
		case "?":
			break
		}
	}
}