package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

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
	ll.deleteAtPosition(2)
	ll.reverseLinkedList()
}

func (ll *Linkedlist) deleteAtPosition(position int) {
	temp := ll.head
	for i := 1; i < position-1; i++ {
		temp = temp.next
	}
	temp.next = temp.next.next
	fmt.Println("Updated Linked list is:")
	ll.printList()

}

func (ll *Linkedlist) deleteAtHead() {
	temp := ll.head
	ll.head = temp.next
	fmt.Println("\nDeleted node at head is: ", temp.data)
	ll.printList()
}

func (ll *Linkedlist) deleteAtTail() {
	temp := ll.head
	for temp.next.next != nil {
		temp = temp.next
	}
	fmt.Println("\nDeleted node at tail is: ", temp.next.data)
	temp.next = nil
	ll.printList()
}

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
	fmt.Println("Reversed a linked list:")
	ll.printList()
	return prev

}

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
