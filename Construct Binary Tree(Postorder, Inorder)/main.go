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
	fmt.Println("Construct Binary Tree from Postorder, Inorder Arrays:")
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	var inorderIndexMap = make(map[int]int)
	for i, v := range inorder {
		inorderIndexMap[v] = i
	}
	postIndex := len(postorder) - 1
	root := constructBinaryTree(postorder, inorderIndexMap, &postIndex, 0, len(inorder)-1)
	fmt.Println("Root of Constructed Binary Tree:", root)
	fmt.Println("Inorder Traversal Of Tree:")
	result := []int{}
	result = inOrderTraversal(root, result)
	fmt.Println(result)
}

func constructBinaryTree(postorder []int, inorderIndexMap map[int]int, postIndex *int, inStart, inEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	rootVal := postorder[*postIndex]
	root := &TreeNode{Val: rootVal}
	*postIndex--
	inOrderIndex := inorderIndexMap[rootVal]
	root.Right = constructBinaryTree(postorder, inorderIndexMap, postIndex, inOrderIndex+1, inEnd)
	root.Left = constructBinaryTree(postorder, inorderIndexMap, postIndex, inStart, inOrderIndex-1)
	return root
}

func inOrderTraversal(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	result = inOrderTraversal(root.Left, result)
	result = append(result, root.Val)
	result = inOrderTraversal(root.Right, result)
	return result
}
