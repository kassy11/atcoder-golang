// 一行の入力をスペース区切りで分割する

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	// スペース区切りで分割
	t := strings.Split(str, " ")
	for i:=0; i<len(t); i++{
		fmt.Println(t[i])
	}
}