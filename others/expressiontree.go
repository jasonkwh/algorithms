package others

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func EvaluateExpressionTree(tree *BinaryTree) int {
	switch tree.Value {
	case -1:
		return EvaluateExpressionTree(tree.Left) + EvaluateExpressionTree(tree.Right)
	case -2:
		return EvaluateExpressionTree(tree.Left) - EvaluateExpressionTree(tree.Right)
	case -3:
		return EvaluateExpressionTree(tree.Left) / EvaluateExpressionTree(tree.Right)
	case -4:
		return EvaluateExpressionTree(tree.Left) * EvaluateExpressionTree(tree.Right)
	default:
		return tree.Value
	}
}
