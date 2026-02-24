package main

// 链表
// 力扣：https://leetcode.cn/problems/linked-list-cycle/

// 方法一：哈希表。遍历所有节点，每次遍历到一个节点时，判断该节点此前是否被访问过。
// func hasCycle(head *ListNode) bool {
// 	seen := map[*ListNode]struct{}{}
// 	for head != nil {
// 		if _, ok := seen[head]; ok {
// 			return true
// 		}
// 		seen[head] = struct{}{}
// 		head = head.Next
// 	}
// 	return false
// }

// 方法二：快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
