package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	result := 0

	for i:=0; i<len(s)-1; i++{
		// iスタートで指定文字以外が出てくるまでループして、
		//lenが最大ならresultに格納
		m:=i
		for (s[m] == 'A' || s[m] == 'C' || s[m] == 'G' || s[m] == 'T') && m < len(s)-1{
			m++
		}
		if result < m-i {
			result = m - i
		}
	}
	fmt.Println(result)
}