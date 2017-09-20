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

func getSubTreeHeight(node *Node) int {
	leftHeight := 0
	rightHeight := 0

	if node == nil {
		return 0
	}

	if node.Left != nil {
		leftHeight = 1 + getSubTreeHeight(node.Left)
	}

	if node.Right != nil {
		rightHeight = 1 + getSubTreeHeight(node.Right)
	}

	if leftHeight > rightHeight {
		return leftHeight
	} else {
		return rightHeight
	}
}

func isNodeBalanced(node *Node) bool {
	leftHeight := getSubTreeHeight(node.Left)
	rightHeight := getSubTreeHeight(node.Right)	
	diff := leftHeight - rightHeight
	
	if diff >= -1 && diff <= 1 {
		return true
	}

	return false	
}

func (t *Tree) IsBalanced() bool {
	if t.Root == nil {
		return true
	}

	leftHeight := isNodeBalanced(t.Root.Left)
	rightHeight := isNodeBalanced(t.Root.Right)

	return leftHeight && rightHeight
}

// NewNode creates a new node
func NewNode(val int) *Node {
	return &Node{Val: val, Left: nil, Right: nil, Parent: nil}
}

func main() {
	fmt.Println("Go Tree")
	tree := &Tree{Root: nil}
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(16)
	tree.Insert(17)
	tree.Insert(44)
	tree.Insert(55)
	tree.printInOrder()
	
	fmt.Printf("Left Subtree height is %d\n", getSubTreeHeight(tree.Root.Left))
	fmt.Printf("Right Subtree height is %d\n", getSubTreeHeight(tree.Root.Right))
	fmt.Printf("Tree balanced? is %v\n", tree.IsBalanced())
}
