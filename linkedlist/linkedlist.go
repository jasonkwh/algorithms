package linkedlist

import "fmt"

type LList struct {
	Value int
	Prev  *LList
	Next  *LList
	Last  *LList
}

func Create(value int) *LList {
	return createLListPrev(value, nil)
}

func createLListPrev(value int, prev *LList) *LList {
	return &LList{
		Value: value,
		Prev:  prev,
	}
}

// O(n) append, deprecated
func (l *LList) AppendOld(value int) {
	// tranverse
	if l.Next != nil {
		l.Next.AppendOld((value))
	}
	l.Next = createLListPrev(value, nil)
}

// O(1)
func (l *LList) Chain(value int) *LList {
	list := createLListPrev(value, l)
	l.Next = list
	return list
}

func (l *LList) Append(value int) {
	if l.Last != nil {
		l.Last = l.Last.Chain(value)
		return
	}
	l.Last = l.Chain(value)
}

func (l *LList) RemoveFirst() {
	if l.Next == nil {
		l = nil
		return
	}
	l.Value = l.Next.Value
	l.Next = l.Next.Next
	l.Next.Last = l.Last
	l.Prev = nil
}

func (l *LList) RemoveLast() {
	l.Last = l.Last.Prev
	l.Last.Next = nil
}

func (l LList) GetFirst() int {
	return l.Value
}

func (l LList) GetLast() int {
	return l.Last.Value
}

func (l LList) Get(index int) int {
	return l.getValue(0, index)
}

func (l LList) getValue(cursor, index int) int {
	if cursor == index {
		return l.Value
	}
	if l.Next != nil {
		cursor++
		return l.Next.getValue(cursor, index)
	}
	return 0
}

func (l LList) String() string {
	if l.Next == nil {
		return fmt.Sprintf("%d", l.Value)
	}
	return fmt.Sprintf("%d -> %s", l.Value, l.Next.String())
}
