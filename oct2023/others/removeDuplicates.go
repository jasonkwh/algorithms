package others

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	linkedList.RemoveDuplicates()
	return linkedList
}

func (l *LinkedList) RemoveDuplicates() {
	l.removeDuplicates(make(map[int]bool), nil)
}

func (l *LinkedList) removeDuplicates(seen map[int]bool, prev *LinkedList) {
	if _, exists := seen[l.Value]; !exists {
		seen[l.Value] = true
	} else {
		prev.Next = l.Next
		if prev.Next != nil {
			prev.Next.removeDuplicates(seen, prev)
			return
		}
	}
	if l.Next != nil {
		l.Next.removeDuplicates(seen, l)
		return
	}
}
