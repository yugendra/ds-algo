package main

import (
	"fmt"
)

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

func (b *BinarySearchTree) DFSInOrder(node *Node, list *[]int) {

	if node.Left != nil {
		b.DFSInOrder(node.Left, list)
	}

	*list = append(*list, node.Value)

	if node.Right != nil {
		b.DFSInOrder(node.Right, list)
	}
}

func (b *BinarySearchTree) DFSPreOrder(node *Node, list *[]int) {

	*list = append(*list, node.Value)

	if node.Left != nil {
		b.DFSPreOrder(node.Left, list)
	}

	if node.Right != nil {
		b.DFSPreOrder(node.Right, list)
	}
}

func (b *BinarySearchTree) DFSPostOrder(node *Node, list *[]int) {

	if node.Left != nil {
		b.DFSPostOrder(node.Left, list)
	}

	if node.Right != nil {
		b.DFSPostOrder(node.Right, list)
	}

	*list = append(*list, node.Value)
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

	var list []int
	b.DFSInOrder(b.Root, &list)
	fmt.Println("Inorder:", list)

	list = []int{}

	b.DFSPreOrder(b.Root, &list)
	fmt.Println("Preorder:", list)

	list = []int{}

	b.DFSPostOrder(b.Root, &list)
	fmt.Println("Postorder:", list)

}
