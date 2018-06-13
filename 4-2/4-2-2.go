package tree

import "fmt"

/*
 * 要改变的内容必须是指针接收者
 * 结构过大也要考虑使用指针接收者，值传递拷贝效率会低
 * 一致性，如有指针接受者，最好都是指针接收者
 */
func (node Node) Print()  { // go 语言传递的都是值
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
