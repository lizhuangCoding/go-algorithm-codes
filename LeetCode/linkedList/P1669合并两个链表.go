package main

// 链表：模拟
// 力扣：https://leetcode.cn/problems/merge-in-between-linked-lists/

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	left, right := &ListNode{}, &ListNode{}
	right = list1
	for i := 0; i < b+1; i++ {
		if i == a-1 {
			left = right
		}
		right = right.Next
	}
	left.Next = list2

	for list2 != nil && list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = right

	return list1
}
