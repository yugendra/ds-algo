package main

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func newLinedList(v int) *linkedList {
	newHead := &node{value: v, next: nil}
	return &linkedList{
		head:   newHead,
		tail:   newHead,
		length: 1,
	}
}

func (ll *linkedList) append(v int) {
	newNode := &node{value: v, next: nil}
	ll.tail.next = newNode
	ll.tail = newNode
	ll.length++
}

func (ll *linkedList) prepend(v int) {
	newNode := &node{value: v, next: ll.head}
	ll.head = newNode
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

	newNode := &node{value: v, next: nil}

	currentNode := ll.traverseToIndex(index - 1)
	tmpNode := currentNode.next
	currentNode.next = newNode
	newNode.next = tmpNode
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

	node := ll.traverseToIndex(index - 1)
	node.next = node.next.next
	ll.length--
}

func (ll *linkedList) Printlist() {
	currentNode := ll.head
	for {
		fmt.Println(currentNode.value)
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
