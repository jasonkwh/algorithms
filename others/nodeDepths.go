package others

func NodeDepths(root *BinaryTree) int {
	// Write your code here.
	return nodeDepths(root, 0)
}

func nodeDepths(bt *BinaryTree, depth int) int {
	output := depth
	depth++
	if bt.Left != nil {
		output += nodeDepths(bt.Left, depth)
	}
	if bt.Right != nil {
		output += nodeDepths(bt.Right, depth)
	}
	return output
}
