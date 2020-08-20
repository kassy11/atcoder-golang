// https://shoman.hatenablog.com/entry/2020/02/25/185456

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stackは[]intのエイリアス
type Stack []string

// Push adds an element
func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

// Pop removes the top element and return it
func (s *Stack) Pop() (string, error) {
	if s.Empty() {
		return "error", fmt.Errorf("stack is empty")
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

// Peek returns the top value
func (s *Stack) Peek() (string, error) {
	if s.Empty() {
		return "error", fmt.Errorf("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

// Size returns the length of stack
func (s *Stack) Size() int {
	return len(*s)
}

// Empty returns true when stack is empty
func (s *Stack) Empty() bool {
	return s.Size() == 0
}

// NewStack generates stack
func NewStack() *Stack {
	s := new(Stack)
	return s
}

func main() {
	stack := NewStack()

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	n, _ := strconv.Atoi(s)

	t := strings.Split(s, " ")
	for i:=0; i<n; i++{
		fmt.Println(t[i])
		stack.Push(t[i])
	}

	fmt.Println(*stack)

	//for i:=0; i<stack.Size(); i++{
	//	switch  {
	//
	//	}
	//}
}