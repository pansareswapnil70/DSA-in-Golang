package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) insert(val int) {
	if val < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: val, Left: nil, Right: nil}
		} else {
			node.Left.insert(val)
		}
	} else if val > node.Value {
		if node.Right == nil {
			node.Right = &Node{Value: val, Left: nil, Right: nil}
		} else {
			node.Right.insert(val)
		}
	}
}

func main() {
	fmt.Println("Validate a Binary Search Tree")
	root := &Node{Value: 2}
	root.insert(1)
	root.insert(3)
	min := math.MinInt64
	max := math.MaxInt64
	boolResult := validateBST(root, min, max)
	fmt.Println("Given Binary Search Tree is valid?: ", boolResult)
}

func validateBST(root *Node, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Value <= min && max <= root.Value {
		return false
	}

	return validateBST(root.Left, min, root.Value) && validateBST(root.Right, root.Value, max)
}
