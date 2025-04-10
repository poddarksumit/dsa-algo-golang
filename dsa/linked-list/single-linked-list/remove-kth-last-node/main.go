package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

func main() {
	fmt.Println("Welcome to Remove Kth node from last!")

	head := node.BuildList([]int{1, 2, 3, 4, 5})
	_, head = RemoveKthFromLastNodeRecursive(head, nil, 0)
	head.PrintList()

	head = node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5})
	head = RemoveKthFromLastNode(head, 0)
	head.PrintList()

	head = node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5})
	head = RemoveKthFromLastNodeV2(head, 0)
	head.PrintList()
}

type Node struct {
	Data int
	Next *Node
}

func RemoveKthFromLastNode(head *node.Node, k int) *node.Node {

	// Base case.
	if head == nil || k < 1 {
		return head
	}

	dummyHead := node.BuildListWithoutPrinting([]int{-1})
	dummyHead.Next = head

	trailer, leader := dummyHead, dummyHead

	for index := 0; index < k && leader != nil; index++ {
		leader = leader.Next
	}

	if leader == nil {
		return head
	}

	for leader.Next != nil {
		leader = leader.Next
		trailer = trailer.Next
	}

	trailer.Next = trailer.Next.Next

	return dummyHead.Next
}

func RemoveKthFromLastNodeV2(head *node.Node, k int) *node.Node {

	// Base case.
	if head == nil || k < 1 {
		return head
	}

	length := findLength(head)

	if length < k {
		return head
	} else if length == k {
		return head.Next
	}

	currentHead := head

	for index := 0; k < length && index < length-k-1; index++ {
		currentHead = currentHead.Next
	}

	currentHead.Next = currentHead.Next.Next

	return head
}

func findLength(head *node.Node) int {
	if head == nil {
		return 0
	}

	count := 0

	for head != nil {
		count++
		head = head.Next
	}
	//fmt.Println("cout: ", count)
	return count
}

// Recursive add to stack so space complixity increases, avoid it.
func RemoveKthFromLastNodeRecursive(head *node.Node, previous *node.Node, k int) (int, *node.Node) {
	if head == nil {
		return 0, head
	}

	lastSize, _ := RemoveKthFromLastNodeRecursive(head.Next, head, k)

	lastSize++

	if k-lastSize == 0 {
		if previous == nil {
			head = head.Next
		} else {
			previous.Next = head.Next
		}
	}

	return lastSize, head
}
