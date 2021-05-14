package main

import "fmt"

type node struct {
	value int
	next  *node
}

type stack struct {
	head   *node
	tail   *node
	length int
}

func newStack() *stack {
	return &stack{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (s *stack) peek() int {
	if s.head == nil {
		fmt.Printf("Stack is empty: ")
		return 0
	}
	return s.head.value
}

func (s *stack) push(v int) {
	newNode := &node{value: v, next: nil}
	if s.head != nil {
		newNode.next = s.head
		s.head = newNode
	} else {
		s.head = newNode
		s.tail = newNode
	}

	s.length++
}

func (s *stack) pop() int {
	if s.head == nil {
		fmt.Print("Stack is empty: ")
		return 0
	}

	if s.head == s.tail {
		s.tail = nil
	}

	topNode := s.head
	s.head = s.head.next
	topNode.next = nil

	s.length--

	return topNode.value
}

func (s *stack) traverseToIndex(index int) *node {
	currentNode := s.head

	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode
}

func main() {
	s := newStack()
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

	fmt.Println(s.head)
	fmt.Println(s.tail)
}
