package main

// 链表
// 力扣：https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// 相似题目：https://leetcode.cn/problems/remove-linked-list-elements/

// func deleteNode(head *ListNode, val int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	if head.Val == val {
// 		return head.Next
// 	}
//
// 	result := head
// 	pre := head // 记录目标节点的前一个节点
// 	head = head.Next
//
// 	for head != nil {
// 		if head.Val == val {
// 			pre.Next = head.Next
// 			break
// 		}
//
// 		head = head.Next
// 		pre = pre.Next
// 	}
//
// 	return result
// }
