package others

func (tree *BST) InOrderTraverse(array []int) []int {
	return tree.inOrderTraverse()
}

func (tree *BST) inOrderTraverse() []int {
	var array []int

	if tree.Left != nil {
		array = append(array, tree.Left.inOrderTraverse()...)
	}
	array = append(array, tree.Value)
	if tree.Right != nil {
		array = append(array, tree.Right.inOrderTraverse()...)
	}

	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	return tree.preOrderTraverse()
}

func (tree *BST) preOrderTraverse() []int {
	var array []int
	array = append(array, tree.Value)

	if tree.Left != nil {
		array = append(array, tree.Left.preOrderTraverse()...)
	}
	if tree.Right != nil {
		array = append(array, tree.Right.preOrderTraverse()...)
	}

	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	return tree.postOrderTraverse()
}

func (tree *BST) postOrderTraverse() []int {
	var array []int

	if tree.Left != nil {
		array = append(array, tree.Left.postOrderTraverse()...)
	}
	if tree.Right != nil {
		array = append(array, tree.Right.postOrderTraverse()...)
	}
	array = append(array, tree.Value)

	return array
}
