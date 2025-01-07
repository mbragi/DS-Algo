package main

import	(
	"fmt"
)
 
type Node struct {
		data any
		next *Node
}

type LinkedList struct {
		head *Node
}

// Function to add a node at the end of the linked list
func (l *LinkedList) addNode(data any) {
	newNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	last := l.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
}

func printList(l *LinkedList) {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}