package main

import "fmt"

func main() {
	cards := make([][]int, 4)
	for i:=0; i<4; i++{
		cards[i] = make([]int, 13)
	}

	var n int
	fmt.Scan(&n)

	for i:=0; i<n; i++{
		var c string
		var x int
		fmt.Scan(&c, &x)
		switch c {
		case "S":
			cards[0][x-1] = 1
		case "H":
			cards[1][x-1] = 1
		case "C":
			cards[2][x-1] = 1
		default:
			cards[3][x-1] = 1
		}

		for i:=0; i<4; i++{
			for j:=0; j<13; j++{
				if cards[i][j] == 0{
					if i == 0{
						fmt.Printf("S %d\n",j+1)
					}else if i == 1{
						fmt.Printf("H %d\n",j+1)
					}else if i == 2{
						fmt.Printf("C %d\n",j+1)
					}else{
						fmt.Printf("D %d\n",j+1)
					}
				}
			}
		}
	}

}
