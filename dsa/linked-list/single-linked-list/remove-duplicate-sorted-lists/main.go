package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

var head *node.Node

func main() {

	fmt.Println("Welcome to Remove Duplicate Nodes from Sorted Single LL.")

	head = node.BuildList(nil)
	removeDuplicateNode(head)

	head = node.BuildList([]int{1})
	head = removeDuplicateNode(head)
	head.PrintList()

	head = node.BuildList([]int{1, 1})
	head = removeDuplicateNode(head)
	head.PrintList()

	head = node.BuildList([]int{1, 5})
	head = removeDuplicateNode(head)
	head.PrintList()

	head = node.BuildList([]int{1, 1, 1, 1, 1, 1, 1, 5, 5})
	head = removeDuplicateNode(head)
	head.PrintList()

	head = node.BuildList([]int{1, 1, 2, 13, 13, 67, 67, 78, 90})
	head = removeDuplicateNode(head)
	head.PrintList()
}

func removeDuplicateNode(head *node.Node) *node.Node {

	if head == nil || head.Next == nil {
		return head
	}

	previous := head
	current := previous.Next

	for current != nil {
		if previous.Value == current.Value {
			previous.Next = current.Next
			current = previous.Next
		} else {
			previous = previous.Next
			current = current.Next
		}
	}

	return head
}
