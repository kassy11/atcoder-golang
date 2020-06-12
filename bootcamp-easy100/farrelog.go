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
	var farelog []int

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
			num, _ := strconv.Atoi(t)
			farelog = append(farelog, num)
		}
	}

	for i:=0; i<len(farelog); i++{
		if point >= float64(farelog[i]){
			point -= float64(farelog[i])
		}else{
			balance -= farelog[i]
			point += float64(farelog[i]) * 0.1
		}
		fmt.Print(balance)
		fmt.Println(point)
	}


}