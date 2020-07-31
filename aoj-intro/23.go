package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var B [4][3][10]int
	for i := 0; i < n; i++ {
		var b, f, r, v int
		fmt.Scanf("%d%d%d%d", &b, &f, &r, &v)
		b--; f--; r--
		B[b][f][r] += v
	}

	for i:= 0; i < 4;i++{
		for j:=0;j<3;j++{
			for k:=0;k < 10;k++{
				fmt.Printf(" %d",B[i][j][k])
			}
			fmt.Print("\n")
		}
		if i != 3{
			fmt.Println("####################")
		}
	}
}
