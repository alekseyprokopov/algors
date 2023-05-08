package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "()[]{}"
	fmt.Println(isValid(input))
}

func isValid(s string) bool {
	data := strings.Split(s, "")
	brackets := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := Stack{}

	for _, item := range data {
		key, ok := brackets[item]
		if !ok {
			stack.Add(item)
		} else {
			if key != stack.Pop() {
				return false
			}
		}
	}

	if !stack.isEmpty() {
		return false
	}
	return true

}

type Stack struct {
	data []string
}

func (s *Stack) Add(bracket string) {
	s.data = append(s.data, bracket)
}

func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}
