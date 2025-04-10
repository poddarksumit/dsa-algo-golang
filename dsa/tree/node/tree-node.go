package node

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func New(data int, leftNode *TreeNode, rightNode *TreeNode) *TreeNode {
	return &TreeNode{
		Value: data,
		Left:  leftNode,
		Right: rightNode,
	}
}

func (node *TreeNode) AddValue(data int, leftNode *TreeNode, rightNode *TreeNode) {

}

func InOrder(node *TreeNode) {

	if node == nil {
		return
	}

	InOrder(node.Left)
	fmt.Println("Data: ", node.Value)
	InOrder(node.Right)
}

func PreOrder(node *TreeNode) {

	if node == nil {
		return
	}

	fmt.Println("Data: ", node.Value)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func PostOrder(node *TreeNode) {

	if node == nil {
		return
	}

	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println("Data: ", node.Value)
}
