package main

// 链表
// 力扣：https://leetcode.cn/problems/swap-nodes-in-pairs/description/

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	pre, cur := dummy, dummy.Next
	// 如果不理解的话，可以画一个链表图来帮助理解
	for cur != nil && cur.Next != nil {
		next := cur.Next
		demo := cur.Next.Next // 记录一下

		// 交换cur和next
		cur.Next = demo
		next.Next = cur
		pre.Next = next

		pre = cur  // 更改pre
		cur = demo // 更改cur
	}
	return dummy.Next
}
