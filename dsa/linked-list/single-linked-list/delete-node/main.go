package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

var head *node.Node

func main() {

	fmt.Println("Welcome to Delete Node in Single LinkedList")

	head = node.BuildList(nil)

	DeleteFromList(head, 9)
	head.PrintList()

	head = node.BuildList([]int{12, 8, 10, 78})

	head = DeleteFromList(head, 6)
	head.PrintList()

	head = DeleteFromList(head, 8)
	head.PrintList()

	head = DeleteFromList(head, 78)
	head.PrintList()

	head = DeleteFromList(head, 10)
	head.PrintList()

	head = DeleteFromList(head, 12)
	head.PrintList()

}

func DeleteFromList(node *node.Node, value int) *node.Node {
	if node == nil {
		return nil
	}

	if node.Value == value {
		//node = node.Next
		return node.Next
	}

	current := node

	for current.Next != nil {

		if current.Next.Value == value {
			current.Next = current.Next.Next
			return node
		}
		current = current.Next
	}

	return node
}
