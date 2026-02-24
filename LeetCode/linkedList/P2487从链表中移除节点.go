package main

// 链表
// 力扣：https://leetcode.cn/problems/remove-nodes-from-linked-list/description/

// 方法一：栈
// func removeNodes(head *ListNode) *ListNode {
// 	var st []*ListNode
// 	for ; head != nil; head = head.Next {
// 		st = append(st, head)
// 	}
// 	for ; len(st) > 0; st = st[:len(st) - 1] {
// 		if head == nil || st[len(st) - 1].Val >= head.Val {
// 			st[len(st) - 1].Next = head
// 			head = st[len(st) - 1]
// 		}
// 	}
// 	return head
// }

// 方法二：先反转链表，然后把链表修改为递增的，然后再反转链表
func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	head = reverseRemoveNodes(head)

	pre := head
	demo := head.Next
	for demo != nil {
		if demo.Val < pre.Val {
			pre.Next = demo.Next
		} else {
			pre = pre.Next
		}

		demo = demo.Next
	}

	// 简化一点：
	// for p := head; p.Next != nil; {
	// 	if p.Val > p.Next.Val {
	// 		p.Next = p.Next.Next
	// 	} else {
	// 		p = p.Next
	// 	}
	// }

	return reverseRemoveNodes(head)
}

func reverseRemoveNodes(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		demo := cur.Next
		cur.Next = pre
		pre = cur
		cur = demo
	}

	return pre
}
