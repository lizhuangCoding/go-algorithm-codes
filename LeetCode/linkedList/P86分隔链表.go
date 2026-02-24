package main

// 链表
// 力扣：https://leetcode.cn/problems/partition-list/description/

// 使用额外空间
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	smallLN, bigLN := &ListNode{}, &ListNode{} // 头部
	smallDemo, bigDemo := smallLN, bigLN

	for head != nil {
		if head.Val < x {
			smallDemo.Next = head
			smallDemo = smallDemo.Next
		} else {
			bigDemo.Next = head
			bigDemo = bigDemo.Next
		}
		head = head.Next
	}
	bigDemo.Next = nil
	smallDemo.Next = bigLN.Next

	return smallLN.Next
}
