package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

var head *node.Node

func main() {

	fmt.Println("Welcome to Check Circular List in Single LinkedList")

	head = node.BuildList(nil)
	fmt.Printf("Is List circular:: %v \n", checkIfLLIsCircular(head))

	head = node.BuildList([]int{12})
	fmt.Printf("Is List circular:: %v \n", checkIfLLIsCircular(head))

	head = node.BuildList([]int{12, 8, 10, 78})
	fmt.Printf("Is List circular:: %v \n", checkIfLLIsCircular(head))

	head = node.BuildList([]int{12, 8, 10, 78, 9})
	fmt.Printf("Is List circular:: %v \n", checkIfLLIsCircular(head))

	var newHeadTemp *node.Node
	newHeadTemp.AddToList(12)
	head.AddToList(110)
	head.AddToList(11)
	newNode := node.New(9)
	head.AddNodeToList(newNode)
	head.AddToList(19)
	head.AddToList(72)
	head.AddNodeToList(newNode)
	fmt.Printf("Is List circular:: %v \n", checkIfLLIsCircular(head))

	var newHead *node.Node
	newNode = node.New(9)
	newHead.AddNodeToList(newNode)
	newHead.AddNodeToList(newNode)
	fmt.Printf("Is List circular:: %v \n", checkIfLLIsCircular(newHead))
}

func checkIfLLIsCircular(node *node.Node) bool {
	if node == nil {
		return false
	}

	slow := node
	fast := node

	for fast != nil &&
		fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
