package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cards := make([]string, n)
	for i:=0; i<len(cards); i++{
		fmt.Scan(&cards[i])
	}

	selection(cards)
	bubble(cards)
}

func selection(nums []string){

	for i := 0; i < len(nums); i++ {
		minj := i
		for j := i; j < len(nums); j++ {
			if nums[j][1:] < nums[minj][1:] {
				minj = j
			}
		}
		if i != minj {
			nums[i], nums[minj] = nums[minj], nums[i]
		}
	}

	for i:=0; i<len(nums); i++{
		fmt.Print(nums[i])
		if i == len(nums) -1{
			fmt.Println("")
		}else{
			fmt.Print(" ")
		}
	}
}

func bubble(nums []string){
	for i:=0; i<len(nums); i++{
		for j:=len(nums)-1; j>i; j--{
			if nums[j-1][1:] > nums[j][1:]{
				tmp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = tmp
			}
		}
	}

	for k:=0; k<len(nums); k++{
		fmt.Print(nums[k])
		if k != len(nums)-1{
			fmt.Print(" ")
		}
	}
	fmt.Println("")
}


