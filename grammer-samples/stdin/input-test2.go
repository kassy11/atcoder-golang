//複数行のデータ入力
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
	a := make([]int, n)
	for i:=0; i<n; i++{
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		n, _ := strconv.Atoi(s)
		a[i] = n
	}
	for i:=0; i<n; i++{
		fmt.Println(a[i])
	}

}