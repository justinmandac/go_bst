package main

import (
	"fmt"
)

// Node tree node
type Node struct {
	Val    int
	Parent *Node
	Left   *Node
	Right  *Node
}

// Tree a binary tree
type Tree struct {
	Root *Node
}

// InOrderTraversal in order traversal
func InOrderTraversal(node *Node) {
	if node.Left != nil {
		InOrderTraversal(node.Left)
	}
	fmt.Println(node.Val)
	if node.Right != nil {
		InOrderTraversal(node.Right)
	}
}

func (t *Tree) printInOrder() {
	InOrderTraversal(t.Root)
}

// Insert inserts into a binary tree
func (t *Tree) Insert(val int) *Tree {
	curr := t.Root

	if curr == nil {
		t.Root = NewNode(val)
		return t
	}

	for curr != nil {
		if val < curr.Val {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr.Left = NewNode(val)
				curr.Left.Parent = curr
				curr = nil
			}
		} else {
			if curr.Right != nil {
				curr = curr.Right
			} else {
				curr.Right = NewNode(val)
				curr.Right.Parent = curr
				curr = nil
			}
		}
	}

	return t
}

// NewNode creates a new node
func NewNode(val int) *Node {
	return &Node{Val: val, Left: nil, Right: nil, Parent: nil}
}

func main() {
	fmt.Println("Go Tree")
	tree := &Tree{Root: nil}
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(5)
	tree.printInOrder()
}
