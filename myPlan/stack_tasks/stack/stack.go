package lifo

import "fmt"

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	up *element
}

func (s *Stack) Add(value, index int) {
	item := &element{
		value: []int{value, index},
		prev:  s.up,
	}
	s.up = item
}

func (s *Stack) Pop() ([]int, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	value := s.up.value
	s.up = s.up.prev
	return value, nil
}

func (s *Stack) IsEmpty() bool {
	return s.up == nil
}
func (s *Stack) Top() ([]int, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.up.value, nil
}

type element struct {
	value []int
	prev  *element
}
