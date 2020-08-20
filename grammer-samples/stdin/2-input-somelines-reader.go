// 複数行の入力を読み込む(paiza 2)
// 読み込み行数も入力で読み込む

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
	str = strings.TrimSpace(str) // 右端の余分な改行文字を削除
	n, _ := strconv.Atoi(str) // 文字列を整数に変換する

	for i:=0; i<n; i++{
		str, _ = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		fmt.Println(str)
	}
}