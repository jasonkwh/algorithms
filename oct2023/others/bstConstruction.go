package others

// func (tree *BST) Insert(value int) *BST {
// 	if value > tree.Value {
// 		if tree.Right != nil {
// 			return tree.Right.Insert(value)
// 		}
// 		tree.Right = &BST{
// 			Value: value,
// 		}
// 	}
// 	if value < tree.Value {
// 		if tree.Left != nil {
// 			return tree.Left.Insert(value)
// 		}
// 		tree.Left = &BST{
// 			Value: value,
// 		}
// 	}
// 	return tree
// }

// func (tree *BST) Contains(value int) bool {
// 	if tree.Value == value {
// 		return true
// 	}
// 	output := false
// 	if tree.Left != nil {
// 		output = tree.Left.Contains(value)
// 	}
// 	if tree.Right != nil {
// 		output = tree.Right.Contains(value)
// 	}
// 	return output
// }

// func (tree *BST) Remove(value int) *BST {
// 	if tree.Value != value {
// 		return tree
// 	}
// 	if tree.Right != nil {
// 		tree.Value = tree.Right.slightlyBigger()
// 	}
// 	return tree
// }

// func (tree *BST) slightlyBigger() int {
// 	if tree.Left != nil {
// 		val := tree.Left.slightlyBigger()
// 		if tree.Left.Value == -1 {
// 			tree.Left = nil
// 		}
// 		return val
// 	}
// 	tempVal := tree.Value
// 	tree.Value = -1
// 	return tempVal
// }
