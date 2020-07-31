package main

import "fmt"

func main() {
	for {
		var H, W int
		fmt.Scanf("%d %d", &H, &W)
		if H == 0 && W == 0{
			break
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if (i+j)%2 == 0{
					fmt.Print("#")
				}else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}
