package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

func main() {

	fmt.Println("Welcome to Print Single LinkedList")

	head := node.BuildList(nil)

	head = node.BuildList([]int{12, 8, 10, 78})

	head.AddToList(90)
	head.AddToList(18)
}
