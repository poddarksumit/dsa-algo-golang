package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

func main() {
	fmt.Println("Welcome to Reverse List")

	head := node.BuildList([]int{1, 2, 3, 4, 5})
	head = ReverseListRecursive(head)
	head.PrintList()

	head = node.BuildList([]int{1, 2, 3, 4, 5})
	reversedHead := RevereListIterative(head)
	reversedHead.PrintList()
}

func ReverseListRecursive(head *node.Node) *node.Node {

	if head == nil || head.Next == nil {
		return head
	}

	value := ReverseListRecursive(head.Next)

	head.Next.Next = head
	head.Next = nil

	return value

}

func RevereListIterative(head *node.Node) *node.Node {

	if head == nil {
		return head
	}

	currentNode := head
	var previousNode *node.Node = nil

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode

		previousNode = currentNode
		currentNode = nextNode
	}

	return previousNode
}
