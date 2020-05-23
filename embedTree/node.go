package embedTree

import (
	"demo-golang/tree"
	"fmt"
)

type Node struct {
	*tree.Node // 内嵌方式扩展
}

func (node *Node) Print() {
	if node == nil || node.Node == nil {
		return
	}

	fmt.Println("这是embedTree.Node的方法")
	node.Node.Print() // 默认取变量的结尾Node
}
