package main

import (
	"fmt"
)

type qNode struct {
	value *Node
	next  *qNode
}

type queue struct {
	first  *qNode
	last   *qNode
	length int
}

func newQueue() *queue {
	return &queue{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

func (q *queue) enqueue(v *Node) {
	newNode := &qNode{value: v, next: nil}
	if q.last == nil {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++
}

func (q *queue) dequeue() *Node {
	if q.length == 0 {
		fmt.Printf("Quest is empty")
		return nil
	}

	n := q.first
	if q.first == q.last {
		q.last = nil
	}
	q.first = q.first.next
	n.next = nil

	q.length--

	return n.value
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (b *BinarySearchTree) insert(value int) {
	newNode := &Node{Value: value}
	if b.Root == nil {
		b.Root = newNode
		return
	}

	cNode := b.Root
	for {
		if value < cNode.Value {
			if cNode.Left == nil {
				cNode.Left = newNode
				break
			} else {
				cNode = cNode.Left
				continue
			}
		} else {
			if cNode.Right == nil {
				cNode.Right = newNode
				break
			} else {
				cNode = cNode.Right
				continue
			}
		}
	}

}

func (b *BinarySearchTree) breadthFirstSearch() []int {

	var bfsList []int

	q := newQueue()
	q.enqueue(b.Root)

	for q.first != nil {
		tNode := q.dequeue()
		bfsList = append(bfsList, tNode.Value)
		if tNode.Left != nil {
			q.enqueue(tNode.Left)
		}
		if tNode.Right != nil {
			q.enqueue(tNode.Right)
		}
	}

	return bfsList
}

func main() {

	b := NewBinarySearchTree()
	b.insert(9)
	b.insert(4)
	b.insert(6)
	b.insert(20)
	b.insert(170)
	b.insert(15)
	b.insert(1)

	bfsArray := b.breadthFirstSearch()
	fmt.Println(bfsArray)

}
