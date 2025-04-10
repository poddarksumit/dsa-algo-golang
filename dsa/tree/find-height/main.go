package main

import (
	"fmt"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/tree/node"
)

func main() {

	fmt.Println("Welcome to find height of Tree")

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

	fmt.Println("Height of this tree:: %v", FindHeight(treehead, 0))

}

func FindHeight(node *node.TreeNode, height int) int {

	if node == nil {
		return height
	}

	leftHeight := FindHeight(node.Left, height+1)
	rightHeight := FindHeight(node.Left, height+1)

	if leftHeight > height {
		height = leftHeight
	}

	if rightHeight > height {
		height = rightHeight
	}

	return height
}
