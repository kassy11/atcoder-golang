// 一行の入力読み込み
// bufio.NewScanner()を使う

package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string

	// Scan()で標準入力を読み込んで、Text()で文字列として取り出す
	if sc.Scan() {
		s = sc.Text()
	}

	fmt.Println(s)
}