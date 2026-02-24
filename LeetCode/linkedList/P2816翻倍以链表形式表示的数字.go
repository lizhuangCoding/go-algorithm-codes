package main

// 链表
// 力扣：https://leetcode.cn/problems/double-a-number-represented-as-a-linked-list/

func doubleIt(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := reverseDoubleIt(head)
	demo := newHead
	jinwei := 0 // 进位
	for demo != nil {
		he := demo.Val*2 + jinwei
		demo.Val = he % 10
		jinwei = he / 10

		if demo.Next == nil && jinwei != 0 {
			demo.Next = &ListNode{Val: jinwei}
			break
		}
		demo = demo.Next
	}

	return reverseDoubleIt(newHead)
}

// 反转链表
func reverseDoubleIt(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		demo := cur.Next
		cur.Next = pre
		pre = cur
		cur = demo
	}
	return pre
}
