package main

import (
	"demo-golang/embedTree"
	"demo-golang/tree"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	node := embedTree.Node{Node: &tree.Node{Value: 1}}
	fmt.Println(node.Value)

	logger, _ := zap.NewProduction()
	logger.Warn("Warning test.")
}
