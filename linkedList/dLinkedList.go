package main

import "fmt"

type node struct {
	value int
	next  *node
	prev  *node
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func newLinedList(v int) *linkedList {
	newHead := &node{value: v, next: nil, prev: nil}
	return &linkedList{
		head:   newHead,
		tail:   newHead,
		length: 1,
	}
}

func (ll *linkedList) append(v int) {
	newNode := &node{value: v, next: nil, prev: ll.tail}
	ll.tail.next = newNode
	ll.tail = newNode
	ll.length++
}

func (ll *linkedList) prepend(v int) {
	newNode := &node{value: v, next: ll.head, prev: nil}
	ll.head = newNode
	newNode.next.prev = newNode
	ll.length++
}

func (ll *linkedList) traverseToIndex(index int) *node {
	currentNode := ll.head

	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode
}

func (ll *linkedList) insert(index, v int) {

	if index == 0 {
		ll.prepend(v)
		return
	}

	if ll.length < index {
		fmt.Println("Invalid index. Index should be less than or equal to", ll.length)
		return
	}

	newNode := &node{value: v, next: nil, prev: nil}

	leader := ll.traverseToIndex(index - 1)
	follower := leader.next

	leader.next = newNode
	newNode.prev = leader

	newNode.next = follower
	follower.prev = newNode

	ll.length++

}

func (ll *linkedList) remove(index int) {
	if index >= ll.length {
		fmt.Println("Invalid index. Index should be less than or equal to", ll.length-1)
		return
	}

	if index == 0 {
		ll.head = ll.head.next
		ll.length--
		return
	}

	leader := ll.traverseToIndex(index - 1)
	follower := leader.next.next

	leader.next = follower
	follower.prev = leader

	ll.length--
}

func (ll *linkedList) Printlist() {
	currentNode := ll.head
	for {
		fmt.Printf("%p-->%d next-->%p prev-->%p\n", currentNode, currentNode.value, currentNode.next, currentNode.prev)
		if currentNode.next == nil {
			break
		}
		currentNode = currentNode.next
	}
}

func main() {
	ll := newLinedList(10)
	ll.append(5)
	ll.append(16)
	ll.prepend(1)
	ll.insert(2, 99)
	ll.Printlist()
	ll.remove(2)
	fmt.Println()
	ll.Printlist()

}
