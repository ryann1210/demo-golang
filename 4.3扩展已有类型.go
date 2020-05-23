package main

import (
	"demo-golang/myTree"
	"demo-golang/queue"
	"demo-golang/tree"
	"fmt"
)

func main() {
	// 组合方式
	node := myTree.Node{Node: &tree.Node{Value: 123}}
	node.Print()

	// 别名方式
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
