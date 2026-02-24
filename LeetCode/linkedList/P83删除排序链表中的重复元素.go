package main

// 链表
// 力扣：https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

// 丑陋的写法：
// func deleteDuplicates(head *ListNode) *ListNode {
// 	dummy := &ListNode{Next: head, Val: -1000}
//
// 	pre := dummy
//
// 	for pre != nil {
// 		cur := pre.Next
//
// 		// demo 记录上一个的val
// 		demo := pre.Val
// 		for cur != nil {
// 			// 出现重复的了
// 			if demo == cur.Val {
// 				cur = cur.Next
// 			} else {
// 				break
// 			}
// 		}
//
// 		pre.Next = cur
// 		pre = pre.Next
// 	}
// 	return dummy.Next
// }

// 一次遍历
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
//
// 	cur := head
// 	for cur.Next != nil {
// 		if cur.Val == cur.Next.Val {
// 			cur.Next = cur.Next.Next
// 		} else {
// 			cur = cur.Next
// 		}
// 	}
//
// 	return head
// }
