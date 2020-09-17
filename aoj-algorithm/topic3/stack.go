
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stackは[]stringのエイリアス
type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (string, error) {
	if s.Empty() {
		return "error", fmt.Errorf("stack is empty")
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

func (s *Stack) Peek() (string, error) {
	if s.Empty() {
		return "error", fmt.Errorf("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func NewStack() *Stack {
	s := new(Stack)
	return s
}

func main() {
	stack := NewStack()

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	str := strings.Split(s, " ")

	for i:=0; i<len(str); i++{
		switch str[i] {
		case "+":
			pop1, _ := stack.Pop()
			pop2, _ := stack.Pop()
			num1, _ := strconv.Atoi(pop1)
			num2, _ := strconv.Atoi(pop2)
			stack.Push(string(num1 + num2))
			break
		case "-":
			pop1, _ := stack.Pop()
			pop2, _ := stack.Pop()
			num1, _ := strconv.Atoi(pop1)
			num2, _ := strconv.Atoi(pop2)
			stack.Push(string(num1 - num2))
			break
		case "*":
			pop1, _ := stack.Pop()
			pop2, _ := stack.Pop()
			num1, _ := strconv.Atoi(pop1)
			num2, _ := strconv.Atoi(pop2)
			fmt.Println(num1, num2)
			stack.Push(string(num1 * num2))
			break
		case "/":
			pop1, _ := stack.Pop()
			pop2, _ := stack.Pop()
			num1, _ := strconv.Atoi(pop1)
			num2, _ := strconv.Atoi(pop2)
			fmt.Println(num1, num2)
			stack.Push(string(num1 / num2))
			break
		default:
			stack.Push(str[i])
			break
		}
	}
	fmt.Println(*stack)
}