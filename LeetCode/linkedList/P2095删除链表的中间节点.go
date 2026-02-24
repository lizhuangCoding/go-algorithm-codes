package main

// 快慢指针
// 力扣：https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/description/

func deleteMiddle(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	slow, fast := head, head
	pre := dummyHead

	for fast != nil && fast.Next != nil {
		pre = slow // pre 记录 slow 的上一个结点
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 这样写诗为了防止出现特殊案例： [1]
	pre.Next = slow.Next // 循环结束后 slow 为待删除节点

	return dummyHead.Next
}
