package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "[([]){}]"
	open := map[string]interface{}{
		"[": nil,
		"(": nil,
		"{": nil,
	}
	close := map[string]interface{}{
		"]": nil,
		")": nil,
		"}": nil,
	}
	dict := map[string]string{
		"]": "[",
		")": "(",
		"}": "{",
	}
	data := strings.Split(input, "")
	stack := Stack{}
	flag := true
	for _, item := range data {
		if _, ok := open[item]; ok {
			stack.Add(item)
		}
		if _, ok := close[item]; ok {
			pair, err := stack.Pop()
			if err != nil {
				flag = false
				break
			}
			if pair != dict[item] {
				flag = false
				break
			}
		}
	}

	if !stack.IsEmpty() {
		flag = false
	}

	fmt.Println(flag)

}

type Stack struct {
	up *element
}

func (s *Stack) Add(value string) {
	item := &element{
		value: value,
		prev:  s.up,
	}
	s.up = item
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("stack is empty")
	}
	value := s.up.value
	s.up = s.up.prev
	return value, nil
}

func (s *Stack) IsEmpty() bool {
	return s.up == nil
}
func (s *Stack) Top() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("stack is empty")
	}
	return s.up.value, nil
}

type element struct {
	value string
	prev  *element
}
