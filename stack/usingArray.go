package main

import "fmt"

const stackSize = 4

type stack struct {
	values []int
}

func (s *stack) peek() int {
	if len(s.values) == 0 {
		fmt.Printf("Stack is empty: ")
		return 0
	}
	return s.values[len(s.values)-1]
}

func (s *stack) pop() int {
	if len(s.values) == 0 {
		fmt.Printf("Stack is empty: ")
		return 0
	}
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}

func (s *stack) push(v int) {
	s.values = append(s.values, v)
}

func main() {
	s := stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)

	fmt.Println(s.peek())
	fmt.Println(s.pop())
	fmt.Println(s.peek())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.peek())
}
