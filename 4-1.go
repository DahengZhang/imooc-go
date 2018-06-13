package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value} // 返回局部变量地址
}

/*
 * 要改变的内容必须是指针接收者
 * 结构过大也要考虑使用指针接收者，值传递拷贝效率会低
 * 一致性，如有指针接受者，最好都是指针接收者
 */
func (node treeNode) print()  { // go 语言传递的都是值
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	fmt.Println(root)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, nil}}
	fmt.Println(nodes)

	root.left.right = createNode(2)
	fmt.Println(root)

	root.print()
	root.setValue(4)
	root.print()

	var nRoot *treeNode
	nRoot.setValue(0)
	nRoot = &root
	nRoot.setValue(9)
	nRoot.print()
	fmt.Println()
	nRoot.traverse()
}
