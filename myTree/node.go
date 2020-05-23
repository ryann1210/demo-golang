package myTree

import (
	"demo-golang/tree"
	"fmt"
)

type Node struct {
	Node *tree.Node
}

func (node *Node) Print() {
	fmt.Println("这是myTree.Node的方法")
	node.Node.Print()
}
