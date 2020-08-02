//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	str := scanner.Text()
//
//	result := make(map[rune]int, 50)
//	for _, s := range str{
//		result[s]++
//	}
//
//	for i, s := range result{
//		fmt.Printf("%s : %d\n", s, i)
//	}
//}