package main

// 链表
// 力扣：https://leetcode.cn/problems/middle-of-the-linked-list/

// 快慢指针
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
