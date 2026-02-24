package main

// 链表
// 力扣：https://leetcode.cn/problems/reverse-nodes-in-k-group/description/

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	// 统计长度
	n := 0
	demo := head
	for demo != nil {
		demo = demo.Next
		n++
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	pre, cur := &ListNode{}, head // pre是当前组的最后一个节点，cur是当前组的下一组的第一个节点

	// k 个一组处理
	for ; n >= k; n -= k {

		for i := 0; i < k; i++ {
			demo := cur.Next
			cur.Next = pre
			pre = cur
			cur = demo
		}

		demo = prev.Next
		prev.Next.Next = cur
		prev.Next = pre
		prev = demo
	}

	return dummy.Next
}
