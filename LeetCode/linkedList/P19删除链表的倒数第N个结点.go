package main

// 链表
// 力扣：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/

// 方法一：计算链表长度
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	length := 0
// 	for ; head != nil; head = head.Next {
// 		length++
// 	}
//
// 	dummy := &ListNode{0, head}
// 	cur := dummy
// 	for i := 0; i < length-n; i++ {
// 		cur = cur.Next
// 	}
// 	cur.Next = cur.Next.Next
// 	return dummy.Next
// }

// 方法二：栈
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	nodes := []*ListNode{}
// 	dummy := &ListNode{0, head}
//
// 	for node := dummy; node != nil; node = node.Next {
// 		nodes = append(nodes, node)
// 	}
// 	prev := nodes[len(nodes)-1-n]
// 	prev.Next = prev.Next.Next
// 	return dummy.Next
// }

// 方法三：双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}

	slow, fast := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}
