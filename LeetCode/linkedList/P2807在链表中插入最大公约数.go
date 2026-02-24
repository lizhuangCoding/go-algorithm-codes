package main

// 链表
// 力扣：https://leetcode.cn/problems/insert-greatest-common-divisors-in-linked-list/description/

// func insertGreatestCommonDivisors(head *ListNode) *ListNode {
// 	pre, cur := head, head.Next
//
// 	for cur != nil {
// 		res := helpInsertGreatestCommonDivisors(pre.Val, cur.Val)
// 		demo := &ListNode{Val: res}
// 		demo.Next = cur
// 		pre.Next = demo
//
// 		pre = cur
// 		cur = cur.Next
// 	}
//
// 	return head
// }
//
// // 求公约数
// func helpInsertGreatestCommonDivisors(a, b int) int {
// 	for i := min(a, b); i >= 1; i-- {
// 		if a%i == 0 && b%i == 0 {
// 			return i
// 		}
// 	}
// 	return 1
// }

// 优化：
func insertGreatestCommonDivisors(head *ListNode) (ans *ListNode) {
	for cur := head; cur.Next != nil; cur = cur.Next.Next {
		cur.Next = &ListNode{gcd(cur.Val, cur.Next.Val), cur.Next}
	}
	return head
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
