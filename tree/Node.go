package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) Print() {
	fmt.Println("Value: ", node.Value)
}
