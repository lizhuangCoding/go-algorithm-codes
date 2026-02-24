package main

// 链表
// 力扣：https://leetcode.cn/problems/odd-even-linked-list/description/

func oddEvenList(head *ListNode) *ListNode {
	n1, n2 := &ListNode{}, &ListNode{}
	demo := n1
	demo2 := n2
	n := 1

	for head != nil {
		if n%2 == 1 {
			n1.Next = &ListNode{Val: head.Val}
			n1 = n1.Next
		} else {
			n2.Next = &ListNode{Val: head.Val}
			n2 = n2.Next
		}

		head = head.Next
		n++
	}

	n1.Next = demo2.Next

	return demo.Next
}

// 优化：
// func oddEvenList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	evenHead := head.Next
// 	odd := head
// 	even := evenHead
// 	for even != nil && even.Next != nil {
// 		odd.Next = even.Next
// 		odd = odd.Next
// 		even.Next = odd.Next
// 		even = even.Next
// 	}
// 	odd.Next = evenHead
// 	return head
// }
