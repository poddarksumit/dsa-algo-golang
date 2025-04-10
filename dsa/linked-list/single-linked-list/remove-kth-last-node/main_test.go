package main

import (
	"testing"

	node "github.com/poddarksumit/dsa-algo-golang/dsa/linked-list/single-linked-list/node"
)

func Test_RemoveKthLastNode(t *testing.T) {

	var tests = []struct {
		testName     string
		inputData    *node.Node
		kthNode      int
		expectedData *node.Node
	}{
		{"Empty node", node.BuildListWithoutPrinting([]int{}), 2, node.BuildListWithoutPrinting([]int{})},
		{"Nil", nil, 2, nil},
		{"Nodes", node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5}), 2, node.BuildListWithoutPrinting([]int{1, 2, 3, 5})},
		{"Nodes head removal", node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5}), 5, node.BuildListWithoutPrinting([]int{2, 3, 4, 5})},
		{"K more  than Nodes", node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5}), 10, node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5})},
		{"K as 0", node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5}), 0, node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5})},
		{"K as -ve", node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5}), -2, node.BuildListWithoutPrinting([]int{1, 2, 3, 4, 5})},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			actualNodes := RemoveKthFromLastNode(test.inputData, test.kthNode)

			for actualNodes != nil {
				if actualNodes.Value != test.expectedData.Value {
					t.Errorf("RemoveKthFromLastNode:: Not match %s, got %v, expected %v", test.testName, actualNodes.Value, test.expectedData.Value)
				}
				actualNodes = actualNodes.Next
				test.expectedData = test.expectedData.Next
			}

		})
	}
}
