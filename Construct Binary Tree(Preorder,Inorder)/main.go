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
	fmt.Println("Construct Binary Tree from Preorder & Inorder arrays:")
	var inOrderMap = make(map[int]int)
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	for i := 0; i < len(inorder); i++ {
		inOrderMap[inorder[i]] = i
	}
	preStart := 0
	root := constructBinaryTree(preorder, inOrderMap, &preStart, 0, len(inorder)-1)
	fmt.Println(root)
	fmt.Println("Inorder Traversal of Tree:")
	result := []int{}
	result = inOrderTraversal(root, result)
	fmt.Println(result)
}

func constructBinaryTree(preorder []int, inOrderMap map[int]int, preStart *int, inStart, inEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	rootVal := preorder[*preStart]
	root := &TreeNode{Val: rootVal}
	*preStart++
	inOrderIndex := inOrderMap[rootVal]

	root.Left = constructBinaryTree(preorder, inOrderMap, preStart, inStart, inOrderIndex-1)
	root.Right = constructBinaryTree(preorder, inOrderMap, preStart, inOrderIndex+1, inEnd)
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
