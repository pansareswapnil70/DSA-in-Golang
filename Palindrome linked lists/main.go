package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func main() {
	fmt.Println("Check if Linked list is Palindrome: ")
	ll := LinkedList{}
	ll.insertAtTail(1)
	ll.insertAtTail(2)
	ll.insertAtTail(3)
	ll.insertAtTail(2)
	ll.insertAtTail(1)
	printList(ll.head)
	result := checkPalindromeUsingSlice(ll.head)
	if result {
		fmt.Println("Given Linked list is Palindrome")
	} else {
		fmt.Println("Given Linked list is not Palindrome")
	}
	res := checkPalindromeUsingReverse(ll.head)
	if res {
		fmt.Println("Given Linked list is Palindrome using reverse")
	} else {
		fmt.Println("Given Linked list is not Palindrome using reverse")
	}
}

//Check linked list is Palindrome using slice/Array.
func checkPalindromeUsingSlice(head *ListNode) bool {
	var data []int
	curr := head
	for ; curr != nil; curr = curr.Next {
		data = append(data, curr.Val)
	}

	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		if data[i] != data[j] {
			return false
		}
	}
	return true
}

//Check linked list is Palindrome by reversing the second half.
func checkPalindromeUsingReverse(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = reverseLinkedList(slow)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		if curr.Next != nil {
			fmt.Printf("%d->", curr.Val)
		} else {
			fmt.Printf("%d\n", curr.Val)
		}
		curr = curr.Next
	}
}

func (ll *LinkedList) insertAtTail(data int) {
	node := &ListNode{data, nil}
	if ll.head == nil {
		ll.head = node
		return
	}
	temp := ll.head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = node
}
