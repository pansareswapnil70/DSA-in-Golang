package main

import "fmt"

type Pair struct {
	TreeNode *TreeNode
	HD       int
}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {
	fmt.Println("Top View of Binary Tree:")
	root := &TreeNode{val: 1}
	root.left = &TreeNode{val: 2}
	root.right = &TreeNode{val: 3}
	root.left.left = &TreeNode{val: 4}
	root.left.right = &TreeNode{val: 5}
	root.right.left = &TreeNode{val: 6}
	root.right.right = &TreeNode{val: 7}

	result := root.topViewOfBinaryTree()
	fmt.Println("Top view of binary tree is: ", result)
}

func (root *TreeNode) topViewOfBinaryTree() []int {
	if root == nil {
		return []int{}
	}

	var HD = make(map[int]int)
	queue := []Pair{}
	queue = append(queue, Pair{root, 0})
	minHD := 0
	maxHD := 0
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		node := front.TreeNode
		hd := front.HD

		if _, exists := HD[hd]; !exists {
			HD[hd] = node.val
		}

		if hd < minHD {
			minHD = hd
		} else if hd > maxHD {
			maxHD = hd
		}

		if node.left != nil {
			queue = append(queue, Pair{TreeNode: node.left, HD: hd - 1})
		}
		if node.right != nil {
			queue = append(queue, Pair{TreeNode: node.right, HD: hd + 1})
		}
	}
	result := []int{}
	for i := minHD; i <= maxHD; i++ {
		result = append(result, HD[i])
	}
	return result
}
