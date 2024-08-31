package main

import (
	"fmt"
)

// struct to define the structure of the node.
type Node struct {
	data int
	next *Node
}

// struct to define the head of the linkedlist.
type Linkedlist struct {
	head   *Node
	length int
}

func main() {
	fmt.Println("Linked List:")
	ll := Linkedlist{}
	ll.insertAtHead(20)
	ll.insertAtHead(10)
	ll.insertAtTail(30)
	ll.insertAtTail(40)
	ll.insertAtPosition(25, 3)
	ll.insertAtPosition(28, 4)
	fmt.Printf("\nLength of Linked list is:  %d", ll.length)
	ll.deleteAtHead()
	ll.deleteAtTail()
	ll.deleteAtPosition(3)
	ll.head = ll.reverseLinkedList()
	ll.head = ll.recursiveReverseLinkedList(ll.head)
	fmt.Println("Reversed a linked list using recursive approach:")
	ll.printList()
}

// Delete a node at specific position in linked list.
func (ll *Linkedlist) deleteAtPosition(position int) {
	temp := ll.head
	for i := 1; i < position-1; i++ {
		temp = temp.next
	}
	temp.next = temp.next.next
	fmt.Println("Updated Linked list after delete at position:", position)
	ll.printList()

}

// Delete a node from head / first node of the linked list.
func (ll *Linkedlist) deleteAtHead() {
	temp := ll.head
	ll.head = temp.next
	fmt.Println("\nDeleted node at head is: ", temp.data)
	fmt.Println("Updated Linked list after delete at head:")
	ll.printList()
}

// Delete a node from tail / last node of the linked list.
func (ll *Linkedlist) deleteAtTail() {
	temp := ll.head
	for temp.next.next != nil {
		temp = temp.next
	}
	fmt.Println("\nDeleted node at tail is: ", temp.next.data)
	fmt.Println("Updated Linked list after delete at tail:")
	temp.next = nil
	ll.printList()
}

// Reverse a linked list using recursive approach.
func (ll *Linkedlist) recursiveReverseLinkedList(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}
	newNode := ll.recursiveReverseLinkedList(node.next)
	node.next.next = node
	node.next = nil

	return newNode
}

// Reverse a linked list using Iterative approach.
func (ll *Linkedlist) reverseLinkedList() *Node {
	var prev *Node
	curr := ll.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	ll.head = prev
	fmt.Println("Reversed a linked list using iterative approach:")
	ll.printList()
	return prev

}

// Print the linked list.
func (ll *Linkedlist) printList() {
	current := ll.head
	for current != nil {
		if current.next != nil {
			fmt.Printf("%d->", current.data)
		} else {
			fmt.Printf("%d", current.data)
		}
		current = current.next
	}
	fmt.Println()
}

// Insert a node at specific position in linked list.
func (ll *Linkedlist) insertAtPosition(data int, position int) {
	var node = &Node{data, nil}
	temp := ll.head
	for i := 1; i < position-1; i++ {
		temp = temp.next
	}
	next := temp.next
	temp.next = node
	node.next = next
	ll.length += 1
	fmt.Println("Node inserted at position:")
	ll.printList()
}

// Insert a node at head / first position in linked list.
func (ll *Linkedlist) insertAtHead(data int) {
	node := &Node{data, nil}
	if ll.head == nil {
		ll.head = node
	} else {
		temp := ll.head
		node.next = temp
		ll.head = node
	}
	ll.length += 1
	fmt.Println("Node inserted at head:")
	ll.printList()
}

// Insert a node at tail / last position in linked list.
func (ll *Linkedlist) insertAtTail(data int) {
	node := &Node{data, nil}
	temp := ll.head
	if ll.head == nil {
		ll.head = node
	} else {
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = node
	}
	ll.length += 1
	fmt.Println("Node inserted at tail:")
	ll.printList()
}
