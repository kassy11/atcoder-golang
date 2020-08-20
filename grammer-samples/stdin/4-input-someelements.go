// 複数要素をスペース区切りで入力する
// 要素数も入力する

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	num, _ := strconv.Atoi(str)

	reader = bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	t := strings.Split(str, " ")

	for i:=0; i<num; i++{
		fmt.Println(t[i])
	}
}