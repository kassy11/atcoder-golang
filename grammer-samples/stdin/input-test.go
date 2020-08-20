// 一行でスペース区切りで読み込む

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
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	n, _ := strconv.Atoi(s)

	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)
	t := strings.Split(s, " ")
	for i:=0; i<n; i++{
		fmt.Println(t[i])
	}

}