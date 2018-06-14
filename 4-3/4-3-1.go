package tree

type Node struct {
	Value int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value} // 返回局部变量地址
}
