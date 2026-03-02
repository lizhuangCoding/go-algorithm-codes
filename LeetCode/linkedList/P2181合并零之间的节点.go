package main

// 链表
// 力扣：https://leetcode.cn/problems/merge-nodes-in-between-zeros/description/

// 在原链表上修改：
// func mergeNodes(head *ListNode) *ListNode {
// 	begin := head
// 	// result :=begin
// 	for begin != nil && begin.Next != nil {
// 		he := 0
// 		demo := begin.Next
// 		for demo != nil && demo.Val != 0 {
// 			he += demo.Val
// 			demo = demo.Next
// 		}
//
// 		begin.Val = he
// 		// 处理尾部
// 		if demo.Val == 0 && demo.Next == nil {
// 			demo = nil
// 		}
// 		begin.Next = demo
// 		begin = begin.Next
// 	}
// 	return head
// }

// 或者：模拟
func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	total := 0
	for cur := head.Next; cur != nil; cur = cur.Next {
		if cur.Val == 0 {
			node := &ListNode{Val: total}
			tail.Next = node
			tail = tail.Next
			total = 0
		} else {
			total += cur.Val
		}
	}

	return dummy.Next
}
