// 半角スペースのある一行を読み込む(paiza 1)
// bufio.NewReader()を利用する

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
}