package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextString() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() (int, error) {
	return strconv.Atoi(nextString())
}

func main() {
	scanner.Split(bufio.ScanWords)

	str := strings.Split(nextString(), "")
	q, _ := nextInt()

	for i := 0; i < q; i++ {
		cmd := nextString()
		a, _ := nextInt()
		b, _ := nextInt()

		// fmt.Println(cmd, a, b, str)

		if cmd == "print" {
			fmt.Println(strings.Join(str[a:(b+1)], ""))
		} else if cmd == "reverse" {
			part := make([]string, b+1-a)
			copy(part, str[a:(b+1)])
			// fmt.Println(part)
			for j := 0; j <= b-a; j++ {
				// fmt.Println(str[a+j], part[(b-a)-j])
				str[a+j] = part[(b-a)-j]
			}
		} else {
			p := strings.Split(nextString(), "")
			for j := a; j <= b; j++ {
				str[j] = p[j-a]
			}
		}
	}
}

