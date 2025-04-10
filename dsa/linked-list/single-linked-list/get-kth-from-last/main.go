package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

var head *node.Node

func main() {

	fmt.Println("Welcome to Get Kth Item from Last in Single LinkedList")

	head = node.BuildList(nil)
	k := 2

	head = node.BuildList([]int{12, 8, 10, 78})

	_, nodeOfKthItem := GetKthItemFromLast(head, 2)
	if nodeOfKthItem == nil {
		fmt.Printf("No item found with this index\n")
	} else {
		fmt.Printf("%v item from last is %v\n", k, nodeOfKthItem.Value)
	}
	nodeOfKthItem = GetKthItemFromLastIterative(head, 2)
	if nodeOfKthItem == nil {
		fmt.Printf("No item found with this index for iterative func\n")
	} else {
		fmt.Printf("%v item from last is %v for iterative func\n", k, nodeOfKthItem.Value)
	}

	_, nodeOfKthItem = GetKthItemFromLast(head, 12)
	if nodeOfKthItem == nil {
		fmt.Printf("No item found with this index\n")
	} else {
		fmt.Printf("%v item from last is %v\n", k, nodeOfKthItem.Value)
	}
	nodeOfKthItem = GetKthItemFromLastIterative(head, 12)
	if nodeOfKthItem == nil {
		fmt.Printf("No item found with this index for iterative func\n")
	} else {
		fmt.Printf("%v item from last is %v for iterative func\n", k, nodeOfKthItem.Value)
	}

	_, nodeOfKthItem = GetKthItemFromLast(head, 4)
	if nodeOfKthItem == nil {
		fmt.Printf("No item found with this index\n")
	} else {
		fmt.Printf("%v item from last is %v\n", k, nodeOfKthItem.Value)
	}
	nodeOfKthItem = GetKthItemFromLastIterative(head, 4)
	if nodeOfKthItem == nil {
		fmt.Printf("No item found with this index for iterative func\n")
	} else {
		fmt.Printf("%v item from last is %v for iterative func\n", k, nodeOfKthItem.Value)
	}

}

func GetKthItemFromLast(node *node.Node, kthItem int) (int, *node.Node) {

	if node == nil {
		return 0, nil
	}

	indexFromLast, nodeOfKthItem := GetKthItemFromLast(node.Next, kthItem)

	if nodeOfKthItem != nil {
		return indexFromLast, nodeOfKthItem
	}

	indexFromLast++

	if kthItem == indexFromLast {
		return indexFromLast, node
	}

	return indexFromLast, nil

}

func GetKthItemFromLastIterative(node *node.Node, k int) *node.Node {

	if node == nil {
		return nil
	}

	slowNode, fastNode := node, node

	for i := 0; i < k && fastNode != nil; i++ {
		fastNode = fastNode.Next
	}

	if fastNode == nil {
		return nil
	}

	for fastNode != nil {
		fastNode = fastNode.Next
		slowNode = slowNode.Next
	}

	return slowNode
}
