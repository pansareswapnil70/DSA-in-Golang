package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}
type Tree struct {
	Root *Node
}

func main() {
	t := Tree{}
	t.insertNode(50)
	t.insertNode(40)
	t.insertNode(60)
	t.insertNode(30)
	inOrderTraversal(t.Root)
	fmt.Println("\nDepth First Traversal:")
	depthFirstTraversal(t.Root)
	fmt.Println("\nBreadth First Traversal:")
	breadthFirstTraversal(t.Root)
}

func (t *Tree) insertNode(val int) {
	if t.Root == nil {
		t.Root = &Node{Val: val}
	} else {
		insert(t.Root, val)
	}
}

func insert(node *Node, val int) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = &Node{Val: val}
		} else {
			insert(node.Left, val)
		}
	} else if val >= node.Val {
		if node.Right == nil {
			node.Right = &Node{Val: val}
		} else {
			insert(node.Right, val)
		}
	}
}

func inOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	inOrderTraversal(node.Left)
	fmt.Printf("%d ", node.Val)
	inOrderTraversal(node.Right)
}

func depthFirstTraversal(n *Node) {
	var stack []*Node
	if n == nil {
		return
	}

	stack = append(stack, n)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
}

func breadthFirstTraversal(n *Node) {
	var queue []*Node
	if n == nil {
		return
	}
	queue = append(queue, n)
	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", front.Val)
		if front.Left != nil {
			queue = append(queue, front.Left)
		}
		if front.Right != nil {
			queue = append(queue, front.Right)
		}
	}
}
