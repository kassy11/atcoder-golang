package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main(){
	var balance int
	var point float64
	var fare []int

	fmt.Print("C ")
	fmt.Scan(&balance)
	b := true
	scanner := bufio.NewScanner(os.Stdin)

	for b {
		fmt.Print("F ")
		b = scanner.Scan()
		t := scanner.Text()
		// endで入力を終了する
		if t == "end" {
			break
		}else{
			scan_num, _ := strconv.Atoi(t)
			fare = append(fare, scan_num)
		}
	}

	for i:=0; i<len(fare); i++{
		if point >= float64(fare[i]){
			point -= float64(fare[i])
		}else{
			balance -= fare[i]
			point += float64(fare[i]) * 0.1
		}
		fmt.Printf("%d ", balance)
		fmt.Println(point)
	}


}