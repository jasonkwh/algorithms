package main

import (
	"algorithms/graphs"
)

func main() {
	node := graphs.CreateNode(1)
	node2 := graphs.CreateNode(2)
	node8 := graphs.CreateNode(8)
	node4 := graphs.CreateNode(4)
	node9 := graphs.CreateNode(9)
	node.Connect(node2, node8, node4, node9)
	node7 := graphs.CreateNode(7)
	node5 := graphs.CreateNode(5)
	node2.Connect(node7, node5)
	node5.Connect(node7)
	node7.Connect(node2)
	node3 := graphs.CreateNode(3)
	node6 := graphs.CreateNode(6)
	node3.Connect(node6)
	node9.Connect(node4)
	node10 := graphs.CreateNode(10)
	node9.Connect(node10)
	node11 := graphs.CreateNode(11)
	node4.Connect(node11)
	node.PrintInDepth()
}
