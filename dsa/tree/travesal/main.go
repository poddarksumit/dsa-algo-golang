package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/tree/node"
)

func main() {

	fmt.Println("Welcome to tree travesal")

	treehead := node.New(1,
		node.New(2,
			node.New(4, nil, nil),
			node.New(5, nil,
				node.New(6, nil, nil))),
		node.New(3,
			nil,
			node.New(7,
				node.New(8, nil, nil),
				node.New(9, nil, nil),
			),
		),
	)

	fmt.Println("Inorder Travesal")
	node.InOrder(treehead)

	fmt.Println("PreOrder Travesal")
	node.PreOrder(treehead)

	fmt.Println("PostOrder Travesal")
	node.PostOrder(treehead)
}
