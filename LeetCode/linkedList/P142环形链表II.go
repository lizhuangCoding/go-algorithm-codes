package main

// 链表
// 力扣：https://leetcode.cn/problems/linked-list-cycle-ii/
// 题解：代码随想录：https://www.programmercarl.com/0142.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8II.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 方法一：快慢指针
// func detectCycle(head *ListNode) *ListNode {
// 	slow, fast := head, head
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 		if slow == fast {
// 			for slow != head {
// 				slow = slow.Next
// 				head = head.Next
// 			}
// 			return head
// 		}
// 	}
// 	return nil
// }

// 方法二：哈希表
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)

	for head != nil {
		if m[head] {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}
