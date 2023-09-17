package graphs

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Nodes []*Node
}

func CreateNode(value int) *Node {
	return &Node{
		Value: value,
		Nodes: make([]*Node, 0),
	}
}

func (n *Node) Connect(nodes ...*Node) {
	n.Nodes = append(n.Nodes, nodes...)
}

func (n *Node) PrintInDepth() {
	n.printInDepth(make([]*Node, 0))
}

func (n *Node) PrintInBreadth() {
	n.printInBreadth(make([]*Node, 0))
}

func (n *Node) printInBreadth(markingNodes []*Node) {
	contain := contains(markingNodes, n)
	markingNodes = append(markingNodes, n)
	if len(n.Nodes) != 0 && !contain {
		for _, node := range n.Nodes {
			fmt.Printf("%d -> %d\n", n.Value, node.Value)
		}
		for _, node := range n.Nodes {
			node.printInBreadth(markingNodes)
		}
	}
}

func (n *Node) printInDepth(markingNodes []*Node) {
	contain := contains(markingNodes, n)
	markingNodes = append(markingNodes, n)
	if len(n.Nodes) != 0 && !contain {
		for _, node := range n.Nodes {
			node.printInDepth(markingNodes)
		}
		return
	}
	result := join(markingNodes, " -> ")
	fmt.Println(result)
}

func contains(nodes []*Node, node *Node) bool {
	for _, n := range nodes {
		if n.Value == node.Value {
			return true
		}
	}
	return false
}

// copied from strings.Join :)
func join(elems []*Node, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%d", elems[0].Value)
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(fmt.Sprintf("%d", elems[i].Value))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(fmt.Sprintf("%d", elems[0].Value))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%d", s.Value))
	}
	return b.String()
}
