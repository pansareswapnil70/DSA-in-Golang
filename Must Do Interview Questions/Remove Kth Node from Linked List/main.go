package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {
	fmt.Println("Remove Nth node from the end:")
	head := &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: &ListNode{Value: 4, Next: &ListNode{Value: 5, Next: nil}}}}}
	fmt.Println("Before Removing Nth Node:")
	printLinkedList(head)
	N := 2
	slow := head
	fast := head
	for N > 0 {
		fast = fast.Next
		N -= 1
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	fmt.Println("\nAfter Removing Nth Node:")
	printLinkedList(head)
}

func printLinkedList(ListNode *ListNode) {
	temp := ListNode
	for temp != nil {
		if temp.Next != nil {
			fmt.Printf("%d->", temp.Value)
		} else {
			fmt.Printf("%d", temp.Value)
		}
		temp = temp.Next
	}
}
