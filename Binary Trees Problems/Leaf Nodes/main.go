package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	fmt.Println("Leaf Nodes Problems:")
	root := &Node{Val: 10}
	root.insert(5)
	root.insert(20)
	root.insert(3)
	root.insert(7)
	root.insert(15)
	root.insert(25)
	leafNodes := []int{}
	count := 0
	sumOfLeaves := 0
	leafnodes, countOfLeaves, sumOfLeaves := leafNodeOperations(root, leafNodes, count, sumOfLeaves)
	fmt.Println("Leaf node are: ", leafnodes)
	fmt.Println("No. of leaf nodes are: ", countOfLeaves)
	fmt.Println("Sum of leaf nodes are: ", sumOfLeaves)
}

func leafNodeOperations(root *Node, leafNodes []int, count int, sumOfLeaves int) ([]int, int, int) {

	if root == nil {
		return leafNodes, count, sumOfLeaves
	}
	leafNodes, count, sumOfLeaves = leafNodeOperations(root.Left, leafNodes, count, sumOfLeaves)
	if root.Left == nil && root.Right == nil {
		count++
		sumOfLeaves += root.Val
		leafNodes = append(leafNodes, root.Val)

	}
	leafNodes, count, sumOfLeaves = leafNodeOperations(root.Right, leafNodes, count, sumOfLeaves)
	return leafNodes, count, sumOfLeaves
}

func (node *Node) insert(val int) {
	newNode := &Node{Val: val}
	if val < node.Val {
		if node.Left == nil {
			node.Left = newNode
		} else {
			node.Left.insert(val)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			node.Right.insert(val)
		}
	}
}
