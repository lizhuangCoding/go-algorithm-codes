package main

// 链表
// 力扣：https://leetcode.cn/problems/add-two-numbers/description/

// 不能把链表1，2转换为数字，然后相加再转换为链表，因为数据范围是链表长度[1,100]，转换为数字太大
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
// 	carry := 0
// 	var p *ListNode
//
// 	for l1 != nil || l2 != nil {
// 		num1, num2 := 0, 0
//
// 		if l1 != nil {
// 			num1 += l1.Val
// 			l1 = l1.Next
// 		}
//
// 		if l2 != nil {
// 			num2 += l2.Val
// 			l2 = l2.Next
// 		}
//
// 		sum := num1 + num2 + carry
// 		sum, carry = sum%10, sum/10
//
// 		if head == nil {
// 			head = &ListNode{Val: sum, Next: nil}
// 			p = head
// 		} else {
// 			p.Next = &ListNode{Val: sum, Next: nil}
// 			p = p.Next
// 		}
// 	}
//
// 	if carry > 0 {
// 		p.Next = &ListNode{Val: carry, Next: nil}
// 	}
// 	return
// }
