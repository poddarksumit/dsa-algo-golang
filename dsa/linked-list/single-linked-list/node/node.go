package node

import "fmt"

type Node struct {
	Next  *Node
	Value int
}

func New(data int) *Node {
	return &Node{
		Next:  nil,
		Value: data,
	}
}

func BuildListWithoutPrinting(linkedListvalues []int) *Node {

	if linkedListvalues == nil {
		return nil
	}

	var head *Node
	var linkedLists *Node

	for _, value := range linkedListvalues {

		if head == nil {
			head = New(value)
			linkedLists = head
		} else {
			linkedLists.Next = New(value)
			linkedLists = linkedLists.Next
		}

	}

	return head
}

func BuildList(linkedListvalues []int) *Node {

	if linkedListvalues == nil {
		return nil
	}

	var head *Node
	var linkedLists *Node

	for _, value := range linkedListvalues {

		if head == nil {
			head = New(value)
			linkedLists = head
		} else {
			linkedLists.Next = New(value)
			linkedLists = linkedLists.Next
		}

	}

	head.PrintList()

	return head
}

func (node *Node) AddToList(value int) {

	if node == nil {
		node = New(value)
	}

	linkedLists := node

	for linkedLists.Next != nil {
		linkedLists = linkedLists.Next
	}

	linkedLists.Next = New(value)
}

func (node *Node) AddNodeToList(newNode *Node) {

	if node == nil {
		node = newNode
		return
	}

	linkedLists := node

	for linkedLists.Next != nil {
		linkedLists = linkedLists.Next
	}

	linkedLists.Next = newNode
}

func (head *Node) PrintList() {

	current := head

	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}

	fmt.Printf("nil \n")
}
