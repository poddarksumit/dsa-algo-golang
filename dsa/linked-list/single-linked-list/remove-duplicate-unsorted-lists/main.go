package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

var head *node.Node

func main() {

	fmt.Println("Welcome to Remove Duplicate Nodes from Unsorted Single LL.")

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

	head = node.BuildList([]int{1, 1, 1, 1, 1, 1, 1, 5})
	head = removeDuplicateNode(head)
	head.PrintList()

	head = node.BuildList([]int{1, 1, 1, 1, 1, 1, 1, 1})
	head = removeDuplicateNode(head)
	head.PrintList()

	head = node.BuildList([]int{1, 5, 1, 1, 1, 1, 1, 5})
	head = removeDuplicateNode(head)
	head.PrintList()

	head = node.BuildList([]int{1, 15, 9, 3, 11, 61, 1, 15})
	head = removeDuplicateNode(head)
	head.PrintList()
}

func removeDuplicateNode(head *node.Node) *node.Node {

	if head == nil || head.Next == nil {
		return head
	}

	previous := head
	current := previous.Next
	valuesMap := map[int]bool{}
	valuesMap[previous.Value] = true

	for current != nil {
		if valuesMap[current.Value] {
			previous.Next = current.Next
			current = previous.Next
		} else {
			valuesMap[current.Value] = true
			current = current.Next
			previous = previous.Next
		}
	}

	return head
}
