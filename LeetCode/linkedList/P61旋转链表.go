package main

// 链表
// 力扣：https://leetcode.cn/problems/rotate-list/description/

// 截取某一段位置拼接到链表的前面
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	// 计算总长度
	l := 0
	demo := head
	for demo != nil {
		l++
		demo = demo.Next
	}

	demo = head
	left := demo
	// 需要向右移动的位置
	k = k % l

	if k == 0 {
		return head
	}

	// 找到断开的位置
	for i := 0; i < l-k-1; i++ {
		demo = demo.Next
	}

	d1 := demo.Next
	demo.Next = nil

	// 找到末尾
	d2 := d1
	for d2 != nil && d2.Next != nil {
		d2 = d2.Next
	}
	// 末尾和头部连接
	d2.Next = left

	return d1
}
