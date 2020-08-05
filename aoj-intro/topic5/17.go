package main

import "fmt"

func main() {
	h := make([]int, 0)
	w := make([]int, 0)

	for{
		var hnum, wnum int
		fmt.Scan(&hnum)
		fmt.Scan(&wnum)
		if hnum == 0 && wnum == 0{
			break
		}
		h = append(h, hnum)
		w = append(w, wnum)
	}

	for i:=0; i<len(h); i++{
		for l:=0; l<h[i]; l++{
			for k:=0; k<w[i]; k++{
				fmt.Print("#")
			}
			fmt.Println()
		}
		fmt.Print("\n")
	}
}