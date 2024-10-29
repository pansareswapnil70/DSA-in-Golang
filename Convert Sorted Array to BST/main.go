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
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println("Convert Sorted Array to Binary Search Tree")
	// Create the balanced BST iteratively
	root := sortedArrayToBST(nums)

	// Print in-order traversal of the tree to check
	fmt.Println("In-order traversal of the BST:")
	inorderTraversal(root)

}

func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		inorderTraversal(root.Right)
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{}
	queue := []struct {
		node  *TreeNode
		start int
		end   int
	}{{root, 0, len(nums) - 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		node := curr.node
		start := curr.start
		end := curr.end

		if start > end {
			continue
		}
		mid := (start + end) / 2
		node.Val = nums[mid]

		if start <= mid-1 {
			node.Left = &TreeNode{}
			queue = append(queue, struct {
				node  *TreeNode
				start int
				end   int
			}{node.Left, start, mid - 1})
		}

		if mid+1 <= end {
			node.Right = &TreeNode{}
			queue = append(queue, struct {
				node  *TreeNode
				start int
				end   int
			}{node.Right, mid + 1, end})
		}

	}
	return root
}
