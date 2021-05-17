package main

import "fmt"

type node struct {
	value int
	next  *node
}

type queue struct {
	first  *node
	last   *node
	length int
}

func newQueue() *queue {
	return &queue{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

func (q *queue) peek() int {
	if q.length == 0 {
		fmt.Printf("Quest is empty")
		return 0
	}

	return q.first.value
}

func (q *queue) enqueue(v int) {
	newNode := &node{value: v, next: nil}
	if q.last == nil {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++
}

func (q *queue) dequeue() int {
	if q.length == 0 {
		fmt.Printf("Quest is empty")
		return 0
	}

	n := q.first
	q.first = q.first.next
	n.next = nil

	q.length--

	return n.value
}

func main() {
	q := newQueue()
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)

	fmt.Println(q.peek())
	fmt.Println(q.dequeue())
	fmt.Println(q.peek())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.peek())
}
