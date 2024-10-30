package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	fmt.Println("Binary Search Tree")
	root := &Node{1, nil, nil}
	root.Left = &Node{2, nil, nil}
	root.Right = &Node{3, nil, nil}
	root.Left.Right = &Node{5, nil, nil}
	fmt.Println("Inorder Traversal of BST:")
	root.inOrderTraversal()
	paths := binaryTreePaths(root)
	fmt.Println("\nIterative Binary Tree Paths:", paths)
	recursivepaths := recursivebinaryTreePaths(root)
	fmt.Println("\nRecursive Binary Tree Paths:", recursivepaths)
}

func recursivebinaryTreePaths(node *Node) []string {
	if node == nil {
		return nil
	}
	var paths []string
	var path string = strconv.Itoa(node.Value)
	DFSRecursiveBinaryTreePaths(node, path, &paths)
	return paths
}

func DFSRecursiveBinaryTreePaths(node *Node, path string, paths *[]string) {
	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, path)
		return
	}
	if node.Left != nil {
		DFSRecursiveBinaryTreePaths(node.Left, path+"->"+strconv.Itoa(node.Left.Value), paths)
	}
	if node.Right != nil {
		DFSRecursiveBinaryTreePaths(node.Right, path+"->"+strconv.Itoa(node.Right.Value), paths)
	}
}

// Iterative Binary Tree Paths.
func binaryTreePaths(root *Node) []string {
	if root == nil {
		return nil
	}
	var paths []string
	stack := []struct {
		node *Node
		path string
	}{{node: root, path: fmt.Sprintf("%d", root.Value)}}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := top.node
		path := top.path

		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
		}

		if node.Right != nil {
			stack = append(stack, struct {
				node *Node
				path string
			}{node: node.Right, path: path + "->" + fmt.Sprintf("%d", node.Right.Value)})
		}

		if node.Left != nil {
			stack = append(stack, struct {
				node *Node
				path string
			}{node: node.Left, path: path + "->" + fmt.Sprintf("%d", node.Left.Value)})
		}
	}

	return paths
}

func (node *Node) inOrderTraversal() {
	if node == nil {
		return
	}
	node.Left.inOrderTraversal()
	fmt.Printf("%d ", node.Value)
	node.Right.inOrderTraversal()
}
