package main

import (
	reverseinteger "algorithms/nov2024/reverse_integer"
	"fmt"
)

func main() {
	// node := graphs.CreateNode(1)
	// node2 := graphs.CreateNode(2)
	// node8 := graphs.CreateNode(8)
	// node4 := graphs.CreateNode(4)
	// node9 := graphs.CreateNode(9)
	// node.Connect(node2, node8, node4, node9)
	// node7 := graphs.CreateNode(7)
	// node5 := graphs.CreateNode(5)
	// node2.Connect(node7, node5)
	// node5.Connect(node7)
	// node7.Connect(node2)
	// node3 := graphs.CreateNode(3)
	// node6 := graphs.CreateNode(6)
	// node3.Connect(node6)
	// node9.Connect(node4)
	// node10 := graphs.CreateNode(10)
	// node9.Connect(node10)
	// node11 := graphs.CreateNode(11)
	// node4.Connect(node11)
	// node.PrintInDepth()
	// node.PrintInBreadth()

	// node1 := addtwonumbers.ListNode{
	// 	Val: 2,
	// 	Next: &addtwonumbers.ListNode{
	// 		Val: 4,
	// 		Next: &addtwonumbers.ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }

	// node2 := addtwonumbers.ListNode{
	// 	Val: 5,
	// 	Next: &addtwonumbers.ListNode{
	// 		Val: 6,
	// 		Next: &addtwonumbers.ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }

	// node1 := addtwonumbers.ListNode{
	// 	Val: 9,
	// 	Next: &addtwonumbers.ListNode{
	// 		Val: 9,
	// 		Next: &addtwonumbers.ListNode{
	// 			Val: 9,
	// 			Next: &addtwonumbers.ListNode{
	// 				Val: 9,
	// 				Next: &addtwonumbers.ListNode{
	// 					Val: 9,
	// 					Next: &addtwonumbers.ListNode{
	// 						Val: 9,
	// 						Next: &addtwonumbers.ListNode{
	// 							Val: 9,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// node2 := addtwonumbers.ListNode{
	// 	Val: 9,
	// 	Next: &addtwonumbers.ListNode{
	// 		Val: 9,
	// 		Next: &addtwonumbers.ListNode{
	// 			Val: 9,
	// 			Next: &addtwonumbers.ListNode{
	// 				Val: 9,
	// 			},
	// 		},
	// 	},
	// }

	test := reverseinteger.Reverse(4321)
	fmt.Println(test)
}
