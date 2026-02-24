package main

// 链表
// 力扣：https://leetcode.cn/problems/merge-two-sorted-lists/description/

// 方法一：循环
// func mergeTwoLists(list1, list2 *ListNode) *ListNode {
// 	dummy := ListNode{} // 用哨兵节点简化代码逻辑
// 	cur := &dummy       // cur 是一个游标指针，负责追踪新链表的尾部
//
// 	for list1 != nil && list2 != nil {
// 		if list1.Val <= list2.Val {
// 			cur.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			cur.Next = list2
// 			list2 = list2.Next
// 		}
// 		cur = cur.Next
// 	}
//
// 	// 拼接剩余链表
// 	if list1 != nil {
// 		cur.Next = list1
// 	}
// 	if list2 != nil {
// 		cur.Next = list2
// 	}
// 	return dummy.Next
// }

// 方法二：递归
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	list := &ListNode{}

	if list1.Val < list2.Val {
		list = list1
		list.Next = mergeTwoLists(list1.Next, list2)
	} else {
		list = list2
		list.Next = mergeTwoLists(list1, list2.Next)
	}

	return list
}
