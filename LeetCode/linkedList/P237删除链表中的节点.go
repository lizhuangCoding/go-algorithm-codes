package main

// 链表
// 力扣：https://leetcode.cn/problems/delete-node-in-a-linked-list/

// 思路：和下一个节点交换。有点类似于脑筋急转弯
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	// 或者直接：
	*node = *node.Next
}
