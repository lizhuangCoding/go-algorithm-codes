package main

// 反转链表中的一段元素
// 力扣：https://leetcode.cn/problems/reverse-linked-list-ii/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head} // 虚拟节点

	// pre0:找到left节点的前一个节点
	pre0 := dummy
	for i := 0; i < left-1; i++ {
		pre0 = pre0.Next
	}

	pre, cur := &ListNode{}, pre0.Next

	// 反转
	for i := 0; i < right-left+1; i++ {
		demo := cur.Next
		cur.Next = pre
		pre = cur
		cur = demo
	}

	// 修改反转过后的链表的头、尾指向
	pre0.Next.Next = cur // pre0.Next 就是初始left位置节点
	pre0.Next = pre
	return dummy.Next
}
