package main

// 链表
// 力扣：https://leetcode.cn/problems/delete-nodes-from-linked-list-present-in-array/description/

func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = 1
	}

	dummy := &ListNode{Next: head}

	pre, cur := dummy, dummy.Next

	for cur != nil {
		if _, ok := m[cur.Val]; ok {
			pre.Next = cur.Next // 删除cur节点
		} else {
			pre = pre.Next
		}

		cur = cur.Next
	}

	return dummy.Next
}
