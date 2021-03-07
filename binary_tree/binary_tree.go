//package binary_tree
package main
import "fmt"

type MyNode struct {
	No int
	Info string
	Left *MyNode
	Right *MyNode
}

func Print(node *MyNode) {
	if node != nil {
		fmt.Println(node.No, node.Info)
		Print(node.Left)
		Print(node.Right)
	}
}

func MakeBinaryTree() *MyNode {
	node1 := &MyNode{No: 1, Info: "cc", Left: nil, Right: nil}
	node2 := &MyNode{No: 2, Info: "cv", Left: nil, Right: nil}
	node3 := &MyNode{No: 3, Info: "cb", Left: nil, Right: nil}
	node4 := &MyNode{No: 4, Info: "cn", Left: nil, Right: nil}
	node5 := &MyNode{No: 5, Info: "cm", Left: nil, Right: nil}
	node6 := &MyNode{No: 6, Info: "cl", Left: nil, Right: nil}
	node7 := &MyNode{No: 7, Info: "ck", Left: nil, Right: nil}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	return node1
}

func main() {
	Root := MakeBinaryTree()
	Print(Root)
}