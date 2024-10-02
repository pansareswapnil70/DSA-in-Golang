package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BST struct {
	root *Node
}

func main() {
	fmt.Println("Binary Search Tree")
	var nodes []int = []int{50, 45, 23, 78, 67, 56, 14}
	t := BST{}
	for _, val := range nodes {
		t.insert(val)
	}
	result := t.inOrderTraversal()
	fmt.Println("Inorder Traversal: ", result)

	result = t.preOrderTraversal()
	fmt.Println("Preorder Traversal: ", result)

	result = t.postOrderTraversal()
	fmt.Println("Postorder Traversal: ", result)

	searchVal := 54
	boolSearch := t.root.Search(searchVal)
	if boolSearch {
		fmt.Printf("%d found in the tree", searchVal)
	} else {
		fmt.Printf("%d not found in the tree", searchVal)
	}
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}
	if val == n.Val {
		return true
	} else if val < n.Val {
		return n.Left.Search(val)
	} else {
		return n.Right.Search(val)
	}
}

func (t *BST) insert(val int) {
	if t.root == nil {
		t.root = &Node{Val: val}
	} else {
		t.root.insertNode(val)
	}
}

func (node *Node) insertNode(val int) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = &Node{Val: val}
		} else {
			node.Left.insertNode(val)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Val: val}
		} else {
			node.Right.insertNode(val)
		}
	}
}

func (t *BST) inOrderTraversal() []int {
	var result []int
	result = t.root.InOrderTraversal(result)
	return result
}

func (node *Node) InOrderTraversal(result []int) []int {
	if node == nil {
		return result
	}
	result = node.Left.InOrderTraversal(result)
	result = append(result, node.Val)
	result = node.Right.InOrderTraversal(result)

	return result
}

func (t *BST) preOrderTraversal() []int {
	var result []int
	if t.root == nil {
		return result
	}
	result = t.root.PreOrderTraversal(result)
	return result
}

func (n *Node) PreOrderTraversal(result []int) []int {
	if n != nil {
		result = append(result, n.Val)
		result = n.Left.PreOrderTraversal(result)
		result = n.Right.PreOrderTraversal(result)
	}
	return result
}

func (t *BST) postOrderTraversal() []int {
	var result []int
	if t.root == nil {
		return result
	}
	result = t.root.PostOrderTraversal(result)
	return result
}

func (n *Node) PostOrderTraversal(result []int) []int {
	if n != nil {
		result = n.Left.PostOrderTraversal(result)
		result = n.Right.PostOrderTraversal(result)
		result = append(result, n.Val)
	}
	return result
}
