package main

// 链表
// 力扣：https://leetcode.cn/problems/swapping-nodes-in-a-linked-list/description/

func swapNodes(head *ListNode, k int) *ListNode {
	n := 0
	demo := head

	for demo != nil {
		demo = demo.Next
		n++
	}

	node1, node2 := head, head
	for i := 0; i < k-1; i++ {
		node1 = node1.Next
	}

	for i := 0; i < n-k; i++ {
		node2 = node2.Next
	}
	node1.Val, node2.Val = node2.Val, node1.Val

	return head
}
