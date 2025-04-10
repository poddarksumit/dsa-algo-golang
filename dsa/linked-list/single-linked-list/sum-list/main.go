package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

func main() {

	fmt.Println("Welcome to Sum of Single LinkedList")

	head := node.BuildList(nil)

	sum := Sum(head)
	fmt.Println("Sum ", sum)
	sum = SumRecurrsive(head)
	fmt.Println("Sum Recurrsive ", sum)

	head = node.BuildList([]int{12, 8, 10, 78})
	sum = Sum(head)
	fmt.Println("Sum ", sum)
	sum = SumRecurrsive(head)
	fmt.Println("Sum Recurrsive ", sum)

	head = node.BuildList([]int{10, 5, 16, 1})
	sum = Sum(head)
	fmt.Println("Sum ", sum)
	sum = SumRecurrsive(head)
	fmt.Println("Sum Recurrsive ", sum)
}

func Sum(head *node.Node) int {
	if head == nil {
		return 0
	}

	current := head
	sum := 0

	for current != nil {
		sum += current.Value
		current = current.Next
	}

	return sum
}

func SumRecurrsive(head *node.Node) int {
	if head == nil {
		return 0
	}

	return head.Value + SumRecurrsive(head.Next)
}
