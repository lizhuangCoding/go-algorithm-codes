package main

// 链表
// 力扣：https://leetcode.cn/problems/palindrome-linked-list/

// 方法一：统一取出链表中的值
// 拼接为字符串：（太慢了）
// func isPalindrome(head *ListNode) bool {
// 	s := ""
// 	for head != nil {
// 		s += strconv.Itoa(head.Val)
// 		head = head.Next
// 	}
//
// 	for left, right := 0, len(s)-1; left <= right; {
// 		if s[left] != s[right] {
// 			return false
// 		}
//
// 		left++
// 		right--
// 	}
//
// 	return true
// }

// 优化：拼接为切片
// func isPalindrome(head *ListNode) bool {
// 	var vals []int
// 	for ; head != nil; head = head.Next {
// 		vals = append(vals, head.Val)
// 	}
//
// 	n := len(vals)
// 	for i, v := range vals[:n/2] {
// 		if v != vals[n-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// 方法二：寻找中间节点+反转链表
// 1.找到前半部分链表的尾节点。
// 2.反转后半部分链表。
// 3.判断是否回文。
// 4.恢复链表。
// 5.返回结果。
func isPalindrome(head *ListNode) bool {
	mid := middleNodeIsPalindrome(head)
	head2 := reverseListIsPalindrome(mid)
	for head2 != nil {
		if head.Val != head2.Val { // 不是回文链表
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

// 找到链表的中间结点
func middleNodeIsPalindrome(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转链表
func reverseListIsPalindrome(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
