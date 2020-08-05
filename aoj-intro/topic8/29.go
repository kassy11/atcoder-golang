// ここからトピック８
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	for _, s := range str{
		if 'a' <= s && s <= 'z'{
			fmt.Print(strings.ToUpper(string(s)))
		}else if 'A' <= s && s <= 'Z'{
			fmt.Print(strings.ToLower(string(s)))
		}else{
			fmt.Print(string(s))
		}
	}
	fmt.Println("")
}