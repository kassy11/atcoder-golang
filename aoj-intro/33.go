// ここからトピック9

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	query := strings.ToLower(nextString())
	count := 0
	for{
		str := nextString()
		if str == "END_OF_TEXT"{
			break
		}

		if strings.ToLower(str) == query{
			count++
		}
	}

	fmt.Println(count)
}