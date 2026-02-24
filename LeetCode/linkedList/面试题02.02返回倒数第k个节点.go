package main

// 链表
// 力扣：https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/

// 快慢指针
func kthToLast(head *ListNode, k int) int {
	// 快指针先走k-1步
	slow, quick := head, head

	for i := 0; i < k-1; i++ {
		quick = quick.Next
	}

	for quick != nil && quick.Next != nil {
		quick = quick.Next
		slow = slow.Next
	}

	return slow.Val
}
