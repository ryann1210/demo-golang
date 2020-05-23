package main

import (
	"demo-golang/embedTree"
	"demo-golang/tree"
	"fmt"
)

func main() {
	node := embedTree.Node{Node: &tree.Node{Value: 1}}
	fmt.Println(node.Value)
}
