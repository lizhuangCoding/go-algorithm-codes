package main

// 链表
// 力扣：https://leetcode.cn/problems/remove-linked-list-elements/

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else { // 如果确认 cur.Next.Val != val 的时候，再向下走，这是为了避免出现：head = [1,2,2,1], val = 2 时候的情况
			cur = cur.Next
		}
	}
	return head
}
