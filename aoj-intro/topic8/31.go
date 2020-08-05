package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	result := make(map[string]int, 50)

	// アルファベットでmapを初期化する
	for i:=0; i<26; i++{
		result[string('a'+i)] = 0
	}

	for{
		str, err := scanner.ReadString('\n')
		if err == io.EOF{
			break
		}
		for _, s := range str{
			result[strings.ToLower(string(s))]++
		}
	}

	for i:=0; i<26; i++{
		fmt.Printf("%s : %d\n", string('a'+i), result[string('a'+i)])
	}
}