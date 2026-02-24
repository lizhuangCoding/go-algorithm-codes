package main

// 链表：反转链表+两数相加
// 力扣：https://leetcode.cn/problems/add-two-numbers-ii/

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	l1 = reverseListAddTwoNumbers(l1)
	l2 = reverseListAddTwoNumbers(l2) // l1 和 l2 反转后，就变成【2. 两数相加】了
	l3 := addTwo(l1, l2)
	return reverseListAddTwoNumbers(l3)
}

func reverseListAddTwoNumbers(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func addTwo(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哨兵节点
	cur := dummy
	carry := 0 // 进位

	for l1 != nil || l2 != nil || carry != 0 { // 有一个不是空节点，或者还有进位，就继续迭代
		if l1 != nil {
			carry += l1.Val // 节点值和进位加在一起
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val // 节点值和进位加在一起
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: carry % 10}
		carry /= 10    // 用于下次循环用的进位
		cur = cur.Next // 下一个节点
	}
	return dummy.Next // 哨兵节点的下一个节点就是头节点
}
