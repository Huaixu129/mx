package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type linkedList struct {
	Head *Node
}

func (list *linkedList) AppendList(date int) {
	newNode := &Node{Data: date, Next: nil}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	lastNode := list.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = newNode
}

func (list *linkedList) InsertAtIndex(data int, index int) {
	newNode := &Node{Data: data, Next: nil}
	if index == 0 {
		newNode.Next = list.Head
		list.Head = newNode
		return
	}
	currentNode := list.Head
	for i := 0; i < index-1 && currentNode != nil; i++ {
		currentNode = currentNode.Next
	}
	if currentNode == nil {
		fmt.Println("Index out of bounds")
		return
	}
	newNode.Next = currentNode.Next
	currentNode.Next = newNode
}

func (list *linkedList) DeleteNode(data int) {
	currentNode := list.Head
	var prevNode *Node = nil
	// Check if the node to be deleted is the head
	if currentNode != nil && currentNode.Data == data {
		list.Head = currentNode.Next
		return
	}
	// Traverse the list to find the node to be deleted
	for currentNode != nil && currentNode.Data != data {
		prevNode = currentNode
		currentNode = currentNode.Next
	}
	// If the node is not found
	if currentNode == nil {
		fmt.Println("Node not found")
		return
	}
	// Update pointers to skip the node to be deleted
	prevNode.Next = currentNode.Next
}

func (list *linkedList) PrintList() {
	currentNode := list.Head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}
func main() {
	mylist := linkedList{}
	mylist.AppendList(0)
	mylist.AppendList(1)
	mylist.AppendList(2)
	mylist.AppendList(3)
	mylist.PrintList()
	fmt.Println()
	mylist.AppendList(4)
	mylist.InsertAtIndex(6, 3)
	mylist.PrintList()
	fmt.Println()
	mylist.DeleteNode(6)
	mylist.PrintList()
}
