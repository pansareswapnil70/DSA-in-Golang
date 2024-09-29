package main

import "fmt"

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

func (node *Node) Search(val int) bool {
	if node == nil {
		return false
	}
	if node.Value == val {
		return true
	} else {
		if val < node.Value {
			return node.Left.Search(val)
		} else {
			return node.Right.Search(val)
		}
	}
}

func (node *Node) inOrderTraversal() {
	if node == nil {
		//fmt.Println("Tree is empty")
		return
	}
	node.Left.inOrderTraversal()
	fmt.Printf("%d ", node.Value)
	node.Right.inOrderTraversal()
}

func (node *Node) preOrderTraversal() {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Value)
	node.Left.preOrderTraversal()
	node.Right.preOrderTraversal()
}

func (node *Node) postOrderTraversal() {
	if node == nil {
		return
	}
	node.Left.postOrderTraversal()
	node.Right.postOrderTraversal()
	fmt.Printf("%d ", node.Value)
}

func main() {
	fmt.Println("Binary Search Tree")
	root := &Node{Value: 25}
	root.insert(20)
	root.insert(22)
	root.insert(5)
	root.insert(10)
	root.insert(12)
	root.insert(28)
	root.insert(30)
	root.insert(36)
	root.insert(38)
	root.insert(40)
	root.insert(48)
	boolSearch := root.Search(38)
	if boolSearch {
		fmt.Println("Node is found")
	} else {
		fmt.Println("Node is not found")
	}
	fmt.Println("Inorder Traversal of BST:")
	root.inOrderTraversal()
	fmt.Println("\nPreorder Traversal of BST:")
	root.preOrderTraversal()
	fmt.Println("\nPostorder Traversal of BST:")
	root.postOrderTraversal()
}
