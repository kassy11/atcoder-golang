// ここからトピック11

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

func stringifyArray(arr []int) string {
	return strings.TrimRight(fmt.Sprintf("%+v", arr)[1:], "]")
}

func swap(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func main() {
	scanner.Split(bufio.ScanWords)
	dice := make([]int, 6)
	for i := 0; i < 6; i++ {
		dice[i], _ = nextInt()
	}
	commands := strings.Split(nextString(), "")
	for i:=range commands{
		cmd := commands[i]

		switch cmd {
		case "S":
			swap(dice, 0, 4)
			swap(dice, 4, 5)
			swap(dice, 5, 1)
		case "E":
			swap(dice, 0, 3)
			swap(dice, 3, 5)
			swap(dice, 5, 2)
		case "N":
			swap(dice, 0, 1)
			swap(dice, 1, 5)
			swap(dice, 5, 4)
		case "W":
			swap(dice, 0, 2)
			swap(dice, 2, 5)
			swap(dice, 5, 3)
		}
	}
	fmt.Println(dice[0])
}

