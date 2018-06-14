package main

import (
	"fmt"
	".."
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	fmt.Println(root)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, nil}}
	fmt.Println(nodes)

	root.Left.Right = tree.CreateNode(2)
	fmt.Println(root)

	root.Print()
	root.SetValue(4)
	root.Print()

	var nRoot *tree.Node
	nRoot.SetValue(0)
	nRoot = &root
	nRoot.SetValue(9)
	nRoot.Print()
	fmt.Println()
	nRoot.Traverse()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
