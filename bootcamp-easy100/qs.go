package main

import "fmt"

func main() {
	var n, a, b int
	var s string
	passer := 0
	foreign_rank := 1
	fmt.Scan(&n, &a, &b)
	fmt.Scan(&s)
	for i:=0; i<len(s); i++{
		if s[i] == 'a' {
			if passer<a+b{
				fmt.Println("Yes")
				passer++
			}else{
				fmt.Println("No")
			}
		}else if s[i] == 'b'{
			if passer<a+b && foreign_rank <= b{
				fmt.Println("Yes")
				passer++
				foreign_rank++
			}else{
				fmt.Println("No")
			}
		}else if s[i] == 'c'{
			fmt.Println("No")
		}
	}

}