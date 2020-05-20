package main

import (
	"fmt"
)

// 1.定义结构体
type treeNode struct {
	value       int
	left, right *treeNode
}

// 2.结构体定义方法
func (node treeNode) print() {
	fmt.Println(node.value)
}

// 3.修改方法 值传递
func (node treeNode) setValue1(value int) {
	node.value = value
}

// 4.修改方法 指针传递
func (node *treeNode) setValue2(value int) {
	node.value = value
}

// 10.遍历
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	// 5.结构体初始化与赋值 若干种形式
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(234)
	root.print() // 3

	root.right.left.setValue1(123)
	root.right.left.print() // 0
	root.right.left.setValue2(123)
	root.right.left.print() // 123

	// 6.指针可以直接调用方法
	pRoot := &root
	pRoot.print()
	pRoot.setValue2(234)
	pRoot.print()

	// 7.创建结构体slice
	nodes := []treeNode{
		{value: 4},
		{},
		{5, nil, &root},
	}
	fmt.Println(nodes)

	// 9.nil调用方法
	//var nilNode *treeNode
	// main方法报错 因为值传递没有值
	//nilNode.setValue1(200)
	// setValue2方法报错 指针传递
	//nilNode.setValue2(200)

	// 10.遍历
	root.traverse()
}

// 8.由于没有构造函数 需要工厂方法控制结构体的创建
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}
