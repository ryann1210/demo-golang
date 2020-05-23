package main

import "demo-golang/tree"

func main() {
	node := tree.Node{Value: 1}
	node.Print()
	node.SetValue(2)
	node.Print()
	node.Left = &tree.Node{Value: 3}
	node.Right = &tree.Node{Value: 4}
	node.Print()
}
