package main

import (
	"encoding/json"
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

func (b *BinarySearchTree) lookup(value int) *Node {
	cNode := b.Root
	for {
		if cNode.Value == value {
			return cNode
		}
		if value < cNode.Value {
			if cNode.Left != nil {
				cNode = cNode.Left
				continue
			} else {
				break
			}
		} else {
			if cNode.Right != nil {
				cNode = cNode.Right
			} else {
				break
			}
		}
	}
	return nil
}

func (b *BinarySearchTree) remove(value int) {

	if b.Root == nil {
		fmt.Println(value, "not found for delete")
		return
	}

	//cNode := b.Root
	parentNode := &Node{}
	leafNode := &Node{}
	rmNode := &Node{}

	if b.Root.Value == value {
		rmNode = b.Root
		if rmNode.Right != nil {
			leafNode = rmNode.Right
			for {
				if leafNode.Left == nil {
					break
				}
				leafNode = leafNode.Left
			}
			leafRightMostNode := leafNode
			for {
				if leafRightMostNode.Right == nil {
					break
				}
				leafRightMostNode = leafRightMostNode.Right
			}
			leafRightMostNode.Right = rmNode.Right
			leafNode.Left = rmNode.Left
			b.Root = leafNode

		} else if rmNode.Left != nil {
			parentNode.Left = rmNode.Left
		} else {
			parentNode.Left = nil
		}
	} else {
		parentNode := b.Root

		for {
			if value < parentNode.Value {
				if parentNode.Left != nil {
					if parentNode.Left.Value == value {
						rmNode = parentNode.Left
						if rmNode.Right != nil {
							leafNode = rmNode.Right
							for {
								if leafNode.Left == nil {
									break
								}
								leafNode = leafNode.Left
							}

							leafRightMostNode := leafNode
							for {
								if leafRightMostNode.Right == nil {
									break
								}
								leafRightMostNode = leafRightMostNode.Right
							}
							if leafRightMostNode != rmNode.Right {
								leafRightMostNode.Right = rmNode.Right
							}
							leafNode.Left = rmNode.Left
							parentNode.Left = leafNode

						} else if rmNode.Left != nil {
							parentNode.Left = rmNode.Left
						} else {
							parentNode.Left = nil
						}
						break
					} else {
						parentNode = parentNode.Left
					}
				} else {
					fmt.Println(value, "not found for delete")
					return
				}
			} else {
				if parentNode.Right != nil {
					if parentNode.Right.Value == value {
						rmNode = parentNode.Right
						if rmNode.Right != nil {
							leafNode = rmNode.Right
							for {
								if leafNode.Left == nil {
									break
								}
								leafNode = leafNode.Left
							}
							leafRightMostNode := leafNode
							for {
								if leafRightMostNode.Right == nil {
									break
								}
								leafRightMostNode = leafRightMostNode.Right
							}
							fmt.Println(leafNode.Value)
							if leafRightMostNode != rmNode.Right {
								leafRightMostNode.Right = rmNode.Right
							}
							leafNode.Left = rmNode.Left
							parentNode.Right = leafNode

						} else if rmNode.Left != nil {
							parentNode.Left = rmNode.Left
						} else {
							parentNode.Left = nil
						}
						break
					} else {
						parentNode = parentNode.Right
					}
				} else {
					fmt.Println(value, "not found for delete")
					return
				}
			}
		}
	}
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

	node := b.lookup(6)
	if node != nil {
		fmt.Println("6 is present")
	} else {
		fmt.Println("6 is not present")
	}

	node = b.lookup(7)
	if node != nil {
		fmt.Println("7 is present")
	} else {
		fmt.Println("7 is not present")
	}

	tree, err := json.Marshal(b)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(tree))

	//b.remove(20)
	b.remove(4)
	tree, err = json.Marshal(b)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(tree))
}
