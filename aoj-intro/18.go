package main

import "fmt"

func main() {
	for {
		var H, W int
		fmt.Scanf("%d %d", &H, &W)
		if H == 0 && W == 0 {
			break
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if i == 0 || j == 0 || i == H-1 || j == W-1 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}
