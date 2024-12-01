package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		l1 = &ListNode{}
	}

	if l2 == nil {
		l2 = &ListNode{}
	}

	val := l1.Val + l2.Val
	tens := int(val / 10)
	if tens != 0 {
		val = val - (10 * tens)

		if l1.Next != nil {
			l1.Next.Val++
		} else {
			l1.Next = &ListNode{
				Val: 1,
			}
		}
	}

	return &ListNode{
		Val:  val,
		Next: AddTwoNumbers(l1.Next, l2.Next),
	}
}
