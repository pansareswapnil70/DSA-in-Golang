package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println("Validate Binary Tree is Heap:")
	root1 := &TreeNode{10,
		&TreeNode{9,
			&TreeNode{7, nil, nil},
			&TreeNode{6, nil, nil},
		},
		&TreeNode{8,
			&TreeNode{5, nil, nil},
			&TreeNode{4, nil, nil},
		},
	}
	root2 := &TreeNode{10,
		&TreeNode{9,
			&TreeNode{7, nil, nil},
			&TreeNode{6, nil, nil},
		},
		&TreeNode{15, nil, nil},
	}
	fmt.Println("Given Binary Tree is Heap for tree rooted at ", root1, "? ", validateBinaryTree(root1))
	fmt.Println("Given Binary Tree is Heap for tree rooted at ", root2, "? ", validateBinaryTree(root2))
}

func validateBinaryTree(node *TreeNode) bool {
	return isCompleteBinaryTree(node) && isBinaryHeap(node)
}

func isCompleteBinaryTree(node *TreeNode) bool {
	if node == nil {
		return true
	}

	queue := []*TreeNode{node}
	var end bool
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if front == nil {
			end = true
		} else {
			if end {
				return false
			}
			queue = append(queue, front.Left, front.Right)
		}
	}
	return true
}

func isBinaryHeap(node *TreeNode) bool {
	if node == nil {
		return true
	}

	if node.Left != nil && node.Left.Val > node.Val {
		return false
	}
	if node.Right != nil && node.Right.Val > node.Val {
		return false
	}

	return isBinaryHeap(node.Left) && isBinaryHeap(node.Right)
}
