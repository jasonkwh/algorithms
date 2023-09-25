package others

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	if len(n.Children) > 0 {
		for i := 0; i < len(n.Children); i++ {
			array = n.Children[i].DepthFirstSearch(array)
		}
	}
	return array
}
