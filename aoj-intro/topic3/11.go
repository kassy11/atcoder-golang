package main

import "fmt"

func main() {
	a := make([]int, 0)
	b := make([]int, 0)
	for{
		var anum, bnum int
		fmt.Scan(&anum)
		fmt.Scan(&bnum)
		if anum == 0 && bnum == 0 {
			break
		}
		a = append(a, anum)
		b = append(b, bnum)
	}

	for i:=0; i<len(a); i++{
		if a[i] >= b[i]{
			fmt.Printf("%d %d\n", b[i], a[i])
		}else{
			fmt.Printf("%d %d\n", a[i], b[i])
		}
	}
}