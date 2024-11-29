package others

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
	return linkedList.GetNode(int(linkedList.Size() / 2))
}

func (l *LinkedList) Size() int {
	return l.size(1)
}

func (l *LinkedList) size(count int) int {
	if l.Next != nil {
		return l.Next.size(count + 1)
	}
	return count
}

func (l *LinkedList) GetNode(index int) *LinkedList {
	return l.getNode(0, index)
}

func (l *LinkedList) getNode(count, index int) *LinkedList {
	if count == index {
		return l
	}
	if l.Next != nil {
		return l.Next.getNode(count+1, index)
	}
	return nil
}
